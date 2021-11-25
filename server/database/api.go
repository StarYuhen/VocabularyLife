package database

import "github.com/gin-gonic/gin"

// Router 数据库绑定路由
func Router(engine *gin.Engine) {
	engine.POST("/api/accountSetup")
}
