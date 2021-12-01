package database

import (
	"VocabularyLife/expend/captcha"
	"VocabularyLife/expend/cryptoapi"
	"github.com/gin-gonic/gin"
)

// Router 数据库路由
func Router(engine *gin.Engine) {
	// 登陆接口--需要验证码--需要加密中间件
	engine.POST("/api/public/major/login", captcha.FuncGinCaptcha(), cryptoapi.FuncGinCrypto(), POSTAccountLogin)
}
