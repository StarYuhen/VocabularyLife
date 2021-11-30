package expend

import (
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/captcha"
	"VocabularyLife/server/cacheRedis"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 杂项---扩展 内容api

// CaptChaNumber 生成数字图形验证码
func CaptChaNumber(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, HttpResult.Success(cat, "请求captcha成功"))
}
