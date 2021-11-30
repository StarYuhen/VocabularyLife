package api

import (
	"VocabularyLife/configs"
	"VocabularyLife/server/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var config = configs.Config.HTTPConfig

// RoutingApi 全局路由启动
func RoutingApi() {
	engine := gin.Default()
	// 定义所有中间件
	engine.Use(cors()) // 跨域中间件

	// 注册接口列表----------------

	// 注册数据库业务接口
	database.Router(engine)

	// 注册列表结束---------------

	// 启动服务端口
	err := engine.Run(":" + config.Port)

	if err != nil {
		logrus.Info("启动gin服务失败,", err)
		return
	}

}
