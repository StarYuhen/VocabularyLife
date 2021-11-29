package database

import "github.com/gin-gonic/gin"

// Router 数据库路由
func Router(engine *gin.Engine) {
	engine.POST("/api/public/major/login", POSTAccountLogin)
}
