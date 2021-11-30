package expend

// 生成图形验证码的返回接口值
type captchaPost struct {
	ID        string `json:"ID"`
	Base64Img string `json:"Base64Img"`
}
