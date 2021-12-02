package database

import (
	"VocabularyLife/expend/captcha"
	"VocabularyLife/expend/cryptoapi"
	"github.com/gin-gonic/gin"
)

// Router 数据库路由
func Router(engine *gin.Engine) {
	// major 重要接口一律需要验证码中间件和加密中间件

	// 登录接口
	engine.POST("/api/public/major/login", captcha.FuncGinCaptcha(), cryptoapi.FuncGinCrypto(), POSTAccountLogin)
	// 注册接口
	engine.POST("/api/public/major/enrollment", captcha.FuncGinCaptcha(), cryptoapi.FuncGinCrypto(), PostAccountEnrollment)
}
