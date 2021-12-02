package flow

import (
	"github.com/sirupsen/logrus"
	"time"
)

// 写入数据库内容

// WriteUserInfo 将登录信息写入数据表
func WriteUserInfo(uid, ip, jwt string) {
	// 创建数据字段
	db.Create(&userinfo{
		uid,
		ip,
		jwt,
	})
}

// WitheAdminAccount 向主用户表插入账号数据
func WitheAdminAccount(uid, user, username, password, imgurl string) {
	// 赋值字段
	var admin adminaccount
	admin.Uid = uid
	admin.User = user
	admin.UserName = username
	admin.Password = password
	admin.ImgUrl = imgurl

	// 默认值传参
	admin.UserLV = 1
	admin.AccountTime = time.Now()
	admin.WordLimit = 600

	logrus.Info("当前注册账号的所有数据:", admin)
	// 创建数据库
	db.Create(&admin)
	logrus.Info("测试是否在这一步")
}
