package flow

// SelectAccount 查询指定用户结果内容 --登录才用这个
func (w *WriteIO) SelectAccount(user string, password string) bool {
	var account adminaccount
	db.Where("user=? and passWord=?", user, password).First(&account)
	w.Admin = account
	return w.Admin.Uid != ""
}

func (w *WriteIO) SelectAccountUser(user string) bool {
	var account adminaccount
	db.Where("user=?", user).First(&account)
	w.Admin = account
	return w.Admin.Uid != ""
}

// ReadUserInfo 查询登录信息表是否有当前字段
func (w *WriteIO) ReadUserInfo() bool {
	var user userinfo
	db.Where("uid=?", w.Admin.Uid).First(&user)
	w.UserInfo = user
	return w.UserInfo.Uid != ""
}
