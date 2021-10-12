package MYSQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func ExecSQLFile()  {
	// 获取当前项目的路径
	PathFile,_:=os.Getwd()
	db, err := sql.Open("mysql", "root:abc123456@tcp(localhost:3306)/VocabularyLife")
	if err != nil {
		logrus.Error("MYSQL 连接失败 ：", err)
	}
	query, err := ioutil.ReadFile(PathFile+"/SQLInit/adminuseraccount.sql")
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(string(query)); err != nil {
		panic(err)
	}
}
