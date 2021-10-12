package MYSQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func MainSqlDB(account string, password string) *sql.DB {
	db, err := sql.Open("mysql", account+":"+password+"@tcp(localhost:3306)/vocabularylife")
	if err != nil {
		logrus.Error("MYSQL 连接失败 ：", err)
	}
	// 设置数据库连接池
	db.SetConnMaxIdleTime(1000)
	return db
}



