package database

// 数据库请求的内容和响应的数据

// Enrollment 请求注册接口列表
type Enrollment struct {
	UserName string `json:"username"`
	Imgurl   string `json:"imgurl"`
}

// PostAccountPassword 请求注册列表
type PostAccountPassword struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
