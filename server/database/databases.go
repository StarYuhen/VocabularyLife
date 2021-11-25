package database

import (
	"../../configs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// 数据库操作

var config = configs.Config.MysqlConfig

var db = sqlInit()

// 初始化数据库配置
func sqlInit() *sql.DB {
	sqlInit := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8", config.UserConfig.User, config.UserConfig.Password, config.UserConfig.Architecture)
	db, err := sql.Open("mysql", sqlInit)
	if err != nil {
		logrus.Error("数据库链接报错：", err)
	}
	// 设置最大连接数 默认为0 也就是没有限制
	db.SetMaxOpenConns(1000)
	// 设置最大空闲连接 每次执行完语句都将连接放入连接池，默认为2
	db.SetConnMaxIdleTime(100)
	return db
}
