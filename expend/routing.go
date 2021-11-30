package expend

import "github.com/gin-gonic/gin"

// 杂项---扩展 路由

func Routing(engine *gin.Engine) {
	// 数字图形验证码
	engine.POST("/api/public/captcha", CaptChaNumber)
}
