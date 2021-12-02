package databasestest

import (
	"VocabularyLife/server/database/flow"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestDatabase(t *testing.T) {
	account, _ := flow.SelectAccount("root", "root")
	logrus.Info("查询指定数据字段的值-->", account)
	logrus.Info("查询是否有指定信息的值", flow.ReadUserInfo(account.Uid))
	flow.WriteUserInfo(account.Uid, "127.0.0.0", "")
}
