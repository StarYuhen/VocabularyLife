package captcha

import (
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/server/cacheRedis"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// RSAcaptcha 生成图形验证码的返回接口值
type RSAcaptcha struct {
	ID     string `json:"id"`
	Number string `json:"number"`
}

// FuncGinCaptcha  gin接口检测验证码函数--用于部分需要验证码的请求接口
func FuncGinCaptcha() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 初始化会话
		session := sessions.Default(ctx)
		var rsacap RSAcaptcha
		// 读取规定的请求头内容
		captcha := ctx.Request.Header.Get("captcha")

		// 解码captcha 数据{id="English",number:"12232"}
		capGet := cryptoapi.RSAGet(captcha)

		errcap := json.Unmarshal(capGet, &rsacap)

		// 获取sessions保存的captcha的id值，做初步检测
		captchaID := session.Get("captcha")

		if errcap != nil {
			logrus.Error("解密失败,图像验证码---", errcap)
			ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("验证码错误，数据不对应"))
			ctx.Abort()
			return
		}

		if captchaID != rsacap.ID {
			logrus.Error("会话保存的验证码id不对应---", rsacap.ID, "----", captchaID)
			ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("验证码ID参数错误"))
			ctx.Abort()
			return
		}

		bo, str := cacheRedis.StringRead(rsacap.ID)

		// 说明验证码不存在或者已过过时
		if !bo || str == "" {
			logrus.Error("查询该验证码不存在---->", rsacap)
			ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("提交的验证码不存在"))
			ctx.Abort()
			return
		}

		// 用于检查验证码id和传入数字是否是正确的

		if rsacap.Number != str {
			logrus.Error("查询该验证码错误---->", rsacap)
			ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("提交的验证码错误"))
			ctx.Abort()
			return
		}

		logrus.Info("验证码校验通过", rsacap)
		ctx.Next()
		// ctx.JSON(http.StatusOK, HttpResult.Success(rsacap, "提交的验证码正确"))

	}
}
