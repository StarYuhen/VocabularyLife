package cryptoapi

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend/HttpResult"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 用于需要加密校验的中间件

// Complex  登录接口传参数据
// CaptchaID     string `json:"captcha"`
// CaptchaNumber string `json:"CaptchaNumber"`
type Complex struct {
	Sign string `json:"sign"`
	Data string `json:"data"`
	Time string `json:"time"`
}

// Times --Complex 解析传参的值
type Times struct {
	Time int `json:"time" mapstructure:"time"`
	Msg  int `json:"msg" mapstructure:"msg"`
}

// Signs --Complex 解析Sign的值
type Signs struct {
	User     string `json:"user" mapstructure:"user"`
	PassWord string `json:"password" mapstructure:"password"`
	Time     int    `json:"time" mapstructure:"time"`
}

// Datas --Complex 解析Data的值
type Datas struct {
	Ip   string `json:"ip"`
	Msg  int    `json:"msg" mapstructure:"msg"`
	Time int    `json:"time" mapstructure:"time"`
}

var config = configs.Config.CryptoConfig

// FuncGinCrypto 用于请求中间件--需要有加密参数的重要接口
/*
 中间件请求参数列表：
sign:加密参数 账号+密码，加密前格式： 账号|密码|time时间戳 {user:账号,password:密码,time:3543645654654} RSA
data:当前设备信息 加密前格式: ip信息|1314|time时间戳 {ip:ip地址,msg:1314，time:3543645654654} RSA
time:时间戳加上（520），加密前格式：{time:当前时间戳，msg:加密额外参数=520} RSA
*/
func FuncGinCrypto() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 接收传参结构体
		var times Times
		var datas Datas
		var signs Signs

		// 获取加密参数头部字段列表
		sign := ctx.Request.Header.Get("Sign")
		data := ctx.Request.Header.Get("Data")
		time := ctx.Request.Header.Get("Time")

		// 解析内容
		timesrc := RSAGet(time)
		datasrc := RSAGet(data)
		signsrc := RSAGet(sign)

		// 判断解析传参内容
		if timesrc == nil || datasrc == nil || signsrc == nil {
			logrus.Error("解析登录加密参数错误")
			ctx.JSON(http.StatusOK, HttpResult.MissingParametersFun("解析登录加密参数错误"))
			ctx.Abort()
			return
		}

		// 将内容赋值给结构体
		errTime := json.Unmarshal(timesrc, &times)
		errSign := json.Unmarshal(signsrc, &signs)
		errData := json.Unmarshal(datasrc, &datas)

		// 当转化结构体错误时
		if errData != nil || errSign != nil || errTime != nil {
			logrus.Error("登录接口将内容赋值给结构体失败-->", errTime, errSign, errData)
			ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("登录接口内部逻辑错误"))
			ctx.Abort()
			return
		}

		// 复杂至极的判断条件
		if times.Time == datas.Time && datas.Time == signs.Time && times.Msg == config.Time && datas.Msg == config.Data {
			// 当这些都满足时说明校验正常
			// 将解析后的值传给context内容，使其读取
			ctx.Set("sign", signs)
			ctx.Set("data", datas)
			ctx.Set("time", times)
			ctx.Next()
			return
		}

		// 说明发生未知错误反应
		logrus.Error("加密中间件接口出现问题----")

		ctx.JSON(http.StatusOK, HttpResult.UnknownErrorFun("加密中间件接口出现问题"))
		ctx.Abort()
		return
	}
}
