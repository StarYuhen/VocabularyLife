package database

import (
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/expend/jwt"
	"github.com/gin-gonic/gin"
)

// Router 数据库路由
func Router(engine *gin.Engine) {
	// major 重要接口一律需要验证码中间件和加密中间件

	// 登录接口 captcha.FuncGinCaptcha(),
	engine.POST("/api/public/major/login", cryptoapi.FuncGinCrypto(), POSTAccountLogin)
	// 注册接口 captcha.FuncGinCaptcha(),
	engine.POST("/api/public/major/enrollment", cryptoapi.FuncGinCrypto(), PostAccountEnrollment)

	// 更新账号头像
	engine.POST("/api/public/account/imgurl", jwt.AuthMiddlewareJWT(), PostAccountUpdateImgurl)
	// 更新账号密码
	engine.POST("/api/public/account/password", jwt.AuthMiddlewareJWT(), PostAccountPassWord)
}
