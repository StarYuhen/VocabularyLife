package api

import (
	"VocabularyLife/server/database"
	"github.com/gin-gonic/gin"
)

// RoutingApi 全局路由启动
func RoutingApi() {
	engine := gin.Default()
	// 定义跨域中间件
	engine.Use(cors())
	engine.POST("/api/login", database.POSTAccountLogin)
}
