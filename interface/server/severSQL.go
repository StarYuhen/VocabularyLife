package server

import (
	"../pkg/MYSQL"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MainSQLServer(run *gin.Engine) {
	// 数据库网关
	DB := MYSQL.MainSqlDB("root", "abc123456")
	run.POST("api/accountSetup", func(context *gin.Context) {
		// 接收数据的结构体
		jsonSetup := make(map[string]interface{})
		context.ShouldBindJSON(&jsonSetup)
		userName := jsonSetup["userName"].(string)
		adminUserAccount := MYSQL.InquireTableAccountForm(DB, userName)
		logrus.Info("POST 请求的所有数据->",jsonSetup)
		context.JSONP(http.StatusOK, adminUserAccount)
	})
}
