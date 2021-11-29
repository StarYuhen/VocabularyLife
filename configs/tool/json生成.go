package main

import (
	"VocabularyLife/expend/cryptoapi"
	"github.com/sirupsen/logrus"
)

func main() {
	// 生成的json必须满足json标准格式
	str := cryptoapi.RSAAdd("{\"time\":3543645654654,\"msg\":520}") // Time的值
	// str := cryptoapi.RSAAdd("{\"user\":\"root\",\"password\" :\"root\",\"time\":3543645654654}") // Sign的值
	// str := cryptoapi.RSAAdd("{\"ip\":\"192.168.0.7\",\"msg\":1314,\"time\":3543645654654}") // Data的值
	logrus.Info("加密后的内容值--->---", str)
	logrus.Info("解密后的内容值---->---", string(cryptoapi.RSAGet(str)))
}
