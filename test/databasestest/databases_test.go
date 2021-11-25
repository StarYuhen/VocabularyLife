package databasestest

import (
	"VocabularyLife/server/database/flow"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestDatabase(t *testing.T) {
	logrus.Info("查询指定数据字段的值-->", flow.SelectAccount(10086))
}
