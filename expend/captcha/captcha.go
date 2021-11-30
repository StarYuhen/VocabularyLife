package captcha

import (
	"VocabularyLife/configs"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

// https://github.com/mojocn/base64Captcha or https://mojotv.cn/go/refactor-base64-captcha

var store = base64Captcha.DefaultMemStore

// 验证码配置
var config = configs.Config.CaptchaConfig

// NumberCaptcha  生成数字图形验证码 --返回验证码对应id和base64编码的图片
func NumberCaptcha() (string, string, string) {
	driver := base64Captcha.NewDriverDigit(config.ImgHeight, config.ImgWidth, config.NumberKeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		logrus.Info("验证码获取失败", err)
		return "", "", ""
	} else {
		// logrus.Info("验证码信息--->", id, b64s)
		number := cp.Store.Get(id, true)
		logrus.Printf("验证码ID:%s,对应数字:%s", id, number)
		return id, b64s, number
	}
}
