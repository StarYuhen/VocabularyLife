package flow

// SelectAccount 查询指定用户结果内容
func SelectAccount(user string, password string) adminaccount {
	var account adminaccount
	db.Where("user=? and passWord=?", user, password).First(&account)
	return account
}

// ReadUserInfo 查询登录信息表是否有当前字段
func ReadUserInfo(uid string) userinfo {
	var user userinfo
	db.Where("uid=?", uid).First(&user)
	return user
}
