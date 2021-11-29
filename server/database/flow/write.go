package flow

// 写入数据库内容

// WriteUserInfo 将登录信息写入数据表
func WriteUserInfo(uid string, ip string, jwt string) {
	// 创建数据字段
	db.Create(userinfo{
		uid,
		ip,
		jwt,
	})
}
