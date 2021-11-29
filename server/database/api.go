package database

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 配置内容信息
var cryptoconfig = configs.Config.CryptoConfig

// POSTAccountLogin  获取当前用户的账号信息，也是登录接口
/*
 请求参数列表：
sign:加密参数 账号+密码，加密前格式： 账号|密码|time时间戳 {user:账号,password:密码,time:3543645654654} RSA
data:当前设备信息 加密前格式: ip信息|1314|time时间戳 {ip:ip地址,msg:1314，time:3543645654654} RSA
time:时间戳加上（520），加密前格式：{time:当前时间戳，msg:加密额外参数=520} RSA
TODO 加密和整体架构待全部完善
*/
func POSTAccountLogin(ctx *gin.Context) {
	var login Login
	var times Times
	var datas Datas
	var signs Signs
	if err := ctx.BindJSON(&login); err != nil {
		logrus.Error("传参解析结构体失败---->", err)
		ctx.JSON(http.StatusOK, server.ParameterErrorFun("发送请求的内容错误"))
		return
	}

	// 将请求的内容进行解密

	// 解析time内容
	time := cryptoapi.RSAGet(login.Time)
	data := cryptoapi.RSAGet(login.Data)
	sign := cryptoapi.RSAGet(login.Sign)

	logrus.Info("请求传入的结果值", string(time), string(data), string(sign))

	// 判断解析传参内容
	if time == nil || data == nil || sign == nil {
		logrus.Error("解析登录加密参数错误")
		ctx.JSON(http.StatusOK, server.MissingParametersFun("解析登录加密参数错误"))
		return
	}

	// 将内容赋值给结构体
	errTime := json.Unmarshal(time, &times)
	errSign := json.Unmarshal(sign, &signs)
	errData := json.Unmarshal(data, &datas)

	if errData != nil || errSign != nil || errTime != nil {
		logrus.Error("登录接口将内容赋值给结构体失败-->", errTime, errSign, errData)
		ctx.JSON(http.StatusOK, server.InternalErrorFun("登录接口内部逻辑错误"))
		return
	}

	// 复杂至极的判断条件
	if times.Time == datas.Time && datas.Time == signs.Time && times.Msg == cryptoconfig.Time && datas.Msg == cryptoconfig.Data {
		// 当这些都满足时说明校监正常
		// TODO 待完成条件都符合的问题
	}

	logrus.Info("具体的值-->", times, datas, signs)

}
