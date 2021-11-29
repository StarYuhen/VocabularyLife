package flow

// UpdateUserInfoIp  更新用户登录信息查询
func UpdateUserInfoIp(uid string, ip string, jwt string) {
	db.Model(&userinfo{}).Where("uid=?", uid).Updates(userinfo{uid, ip, jwt})
}
