package flow

import (
	"VocabularyLife/configs"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库操作 gorm

var config = configs.Config.MysqlConfig

var db = sqlInit()

// 初始化数据库配置
func sqlInit() *gorm.DB {
	sqlInit := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?utf8mb4&parseTime=True&loc=Local", config.UserConfig.User, config.UserConfig.Password, config.UserConfig.Architecture)
	SQL, err := gorm.Open(mysql.Open(sqlInit), &gorm.Config{})
	if err != nil {
		logrus.Error("数据库链接报错：", err)
	}
	// 利用database/sql设置数据库连接池
	sql, err := SQL.DB()

	// 设置最大连接数 默认为0 也就是没有限制
	sql.SetMaxOpenConns(0)
	// 设置最大空闲连接 每次执行完语句都将连接放入连接池，默认为2
	sql.SetConnMaxIdleTime(100)
	return SQL
}
