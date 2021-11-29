package cryptoapi

import (
	"VocabularyLife/configs"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"github.com/sirupsen/logrus"
)

// PrivateKey 恒久不变的私钥内容
var PrivateKey = configs.PrivateKey

// PublicKey 公钥内容
var PublicKey = configs.PublicKey

// RSAGet  提供接口的解密操作
func RSAGet(res string) []byte {
	// 进行base解码
	data, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		logrus.Error("base64解码错误---->", err)
		return nil
	}
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, data)
	if err != nil {
		logrus.Error("rsa解密内容出错---->", err)
		return nil
	}
	return decryptedBytes
}

// RSAAdd 提供内容的加密操作
func RSAAdd(res string) string {
	plain := []byte(res)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, PublicKey, plain)
	if err != nil {
		logrus.Error("rsa加密内容出错---->", err)
		return ""
	}
	// 进行base64编码
	return base64.StdEncoding.EncodeToString(cipherText)
}

// 将解密后的内容赋值给结构体的部分参数
