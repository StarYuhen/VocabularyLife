package api

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend"
	"VocabularyLife/server/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var config = configs.Config

// RoutingApi 全局路由启动
func RoutingApi() {
	engine := gin.Default() // 启动gin默认引擎
	engineCookie := cookie.NewStore([]byte(config.CryptoConfig.Cookie))

	// TODO 中间件注册列表----------------
	// 跨域中间件
	engine.Use(cors())

	// cookie和sessions中间件
	engine.Use(sessions.Sessions(config.CryptoConfig.CookieName, engineCookie))
	// 中间件注册列表结束----------------

	// TODO 注册接口列表----------------

	// 注册数据库业务接口
	database.Router(engine)

	// 杂项数据接口
	expend.Routing(engine)

	// 注册接口列表结束----------------

	// 启动服务端口
	err := engine.Run(":" + config.HTTPConfig.Port)

	if err != nil {
		logrus.Info("启动gin服务失败,", err)
		return
	}

}
