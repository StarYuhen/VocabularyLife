package flow

// UpdateUserInfoIp  更新用户登录信息查询
func UpdateUserInfoIp(uid string, ip string, jwt string) {
	db.Scan(userinfo{
		uid,
		ip,
		jwt,
	})
}
