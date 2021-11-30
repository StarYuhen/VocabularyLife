package database

// 数据库请求的内容和响应的数据

// Login 登录接口传参数据
// CaptchaID     string `json:"captcha"`
// CaptchaNumber string `json:"CaptchaNumber"`
type Login struct {
	Sign string `json:"sign"`
	Data string `json:"data"`
	Time string `json:"time"`
}

// Times --Login 解析传参的值
type Times struct {
	Time int `json:"time" mapstructure:"time"`
	Msg  int `json:"msg" mapstructure:"msg"`
}

// Signs --Login 解析Sign的值
type Signs struct {
	User     string `json:"user" mapstructure:"user"`
	PassWord string `json:"password" mapstructure:"password"`
	Time     int    `json:"time" mapstructure:"time"`
}

// Datas --Login 解析Data的值
type Datas struct {
	Ip   string `json:"ip"`
	Msg  int    `json:"msg" mapstructure:"msg"`
	Time int    `json:"time" mapstructure:"time"`
}
