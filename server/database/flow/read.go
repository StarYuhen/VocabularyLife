package flow

// SelectAccount 查询指定用户结果内容
func SelectAccount(uid int) adminaccount {
	var account adminaccount
	db.Where("uid=?", uid).First(&account)

	return account
}
