package expend

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/captcha"
	"VocabularyLife/server/cacheRedis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 杂项---扩展 内容api

var config = configs.Config.CryptoConfig

// CaptChaNumber 生成数字图形验证码
func CaptChaNumber(ctx *gin.Context) {
	// 初始化sessions对象
	session := sessions.Default(ctx)

	// 随机生成一个图片
	id, base64Img, number := captcha.NumberCaptcha()
	// 由于该扩展库似乎自带的验证方法不对 --自行将id绑定验证码数字提交redis，并设置过期时间
	bo := cacheRedis.StringSet(id, number)
	if !bo {
		// 写入redis失败，自行返回错误信息
		ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("captcha写入redis失败"))
		return
	}

	var cat captchaPost
	cat.Base64Img = base64Img
	cat.ID = id

	// 获取验证码都需要写入sessions --用于为账号绑定

	// 无论如何访问此接口都会刷新session值
	// 插入sessions值
	session.Set(config.CookieList[0], id)
	// 保存sessions值
	err := session.Save()
	if err != nil {
		logrus.Info("保存sessions失败--->", err)
		ctx.JSON(http.StatusOK, HttpResult.UnknownErrorFun("保存验证码数据失败，请刷新"))
		return
	}
	logrus.Info("保存的captcha的id--->", id)

	ctx.JSON(http.StatusOK, HttpResult.Success(cat, "请求captcha成功"))
}
