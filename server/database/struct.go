package database

// 数据库请求的内容和响应的数据

// Login 登录接口传参数据
type Login struct {
	Sign []byte `json:"sign"`
	Data []byte `json:"data"`
	Time []byte `json:"time"`
}

// Times 解析传参的值
var Times struct {
	Time uint64 `json:"time" mapstructure:"time"`
	Msg  int    `json:"msg" `
}
