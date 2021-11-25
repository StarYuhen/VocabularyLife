package database

import (
	"VocabularyLife/api"
	"VocabularyLife/expend/cryptoapi"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// POSTAccountLogin  获取当前用户的账号信息，也是登录接口
/*
 请求参数列表：

sign:加密参数 账号+密码，加密前格式： 账号|密码|time内容 RSA
data:当前设备信息 加密前格式: 设备信息|1314|time内容 RSA
time:时间戳加上（520），加密前格式：{time:当前时间戳，msg:加密额外参数=520} RSA



--- 请在axios配置拦截器，这个加密属于中间件解析的内容

TODO 加密和整体架构待全部完善
*/
func POSTAccountLogin(ctx *gin.Context) {
	var login Login
	if err := ctx.BindJSON(&login); err != nil {
		logrus.Error("传参解析结构体失败---->", err)
		ctx.JSON(http.StatusOK, api.ParameterErrorFun("发送请求的内容错误"))
		return
	}

	// 将请求的内容进行解密

	// 解析time内容
	if time := cryptoapi.RSAGet(login.Time); time == "" {
		logrus.Error("解析time值为空")
		ctx.JSON(http.StatusOK, api.MissingParametersFun("请求的time转换结果为空"))
		return
	}

	logrus.Info(login)
}
