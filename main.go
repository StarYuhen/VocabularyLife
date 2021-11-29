package main

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend/cryptoapi"
	"github.com/sirupsen/logrus"
)

func main() {
	// 界面运行展示配置
	logrus.Info("当前VocabularyLife的版本是:", configs.Config.AppVersion)
	// str:=cryptoapi.RSAAdd("{time:3543645654654,msg:520}")
	str := "XNuJrRWZe4KQkNW6f3ujzSRmAIWqXMe8QVY4tITJYSpwH01Xhn0BVfcUtegxWK6/YrUEIyQD6UInqxhSzLm4RD2+qbe65tOSAQqtH/6rwe/R8HyzYzGKDPx7cziOcS46l/b6iBcV05fokFywNLctEvYn1pUIUxPqiBi6Se5LhjX+z7SxHIp/X8GqetEmOcsq1IdUUzud7DCeU7jAI12PqNSihpLaanklLUBJxlxp3jhA8yNwJ8rR5TwC/2vT8y3/c43J2xFRN02QUb6sDPQkyAXDTT6iF3Q8l69vHfu3951wJc2tmQCj7LpzFfTF2ZLmRiWfjy0duwQT+V1hRBWFrQ=="
	logrus.Info("加密后的内容值--->---", str)
	logrus.Info("解密后的内容值---->---", cryptoapi.RSAGet(str))
}
