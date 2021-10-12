package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

// InquireTableAccountForm 返回用户数据库的所有字段信息
func InquireTableAccountForm(db *sql.DB, UID string) account {
	res, err := db.Query("select * from adminuseraccount  where uid=?", UID)
	if err != nil {
		logrus.Error("数据库查询内容失败user_main:", err)
	}

	var userAccount account

	for res.Next() {
		res.Scan(&userAccount.Uid, &userAccount.UserName,  &userAccount.UserLV, &userAccount.ImgUrl, &userAccount.ContinuousCheckIn, &userAccount.AccumulativeCheckIn, &userAccount.MaximumContinuous)
	}
	logrus.Info("查询的是指定内容数据：", userAccount)
	return userAccount
}
