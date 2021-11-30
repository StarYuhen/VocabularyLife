package database

import (
	"VocabularyLife/expend/captcha"
	"github.com/gin-gonic/gin"
)

// Router 数据库路由
func Router(engine *gin.Engine) {
	// 登陆接口--需要验证码正确
	engine.POST("/api/public/major/login", captcha.FuncGinCaptcha(), POSTAccountLogin)
}
