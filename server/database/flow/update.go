package flow

import "github.com/sirupsen/logrus"

// UpdateUserInfoIp  更新用户登录信息查询
func (w *WriteIO) UpdateUserInfoIp(ip, jwt string) {
	var user userinfo
	user.Uid = w.Admin.User
	user.Ip = ip
	user.Jwt = jwt
	err := db.Model(&userinfo{}).Where("uid=?", w.Admin.Uid).Updates(user).Error
	logrus.Info("更新用户登录信息查询-->", err)
	if err == nil {
		w.UserInfo = user
	}
}

// UpdateAdminAccountImgurl  更新账号头像
func (w *WriteIO) UpdateAdminAccountImgurl(imgurl string) {
	var admin adminaccount
	admin.User = w.Admin.User
	err := db.Model(&admin).Select("imgUrl").Updates(map[string]string{"imgUrl": imgurl}).Error
	logrus.Info("更新账号头像-->", err)
	if err == nil {
		w.Admin.ImgUrl = imgurl
	}
}

// UpdateAdminAccountPassword 更新主账号密码
func (w *WriteIO) UpdateAdminAccountPassword(password string) {
	var admin adminaccount
	admin.User = w.Admin.User
	err := db.Model(&admin).Select("passWord").Updates(map[string]interface{}{"passWord": password}).Error
	logrus.Info("更新主账号密码-->", err)
	if err == nil {
		w.Admin.Password = password
	}
}
