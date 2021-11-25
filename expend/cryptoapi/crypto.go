package cryptoapi

import (
	"VocabularyLife/configs"
	"crypto/rand"
	"crypto/rsa"
	"github.com/sirupsen/logrus"
)

// PrivateKey 恒久不变的私钥内容
var PrivateKey = configs.PrivateKey

// PublicKey 公钥内容
var PublicKey = configs.PublicKey

// RSAGet  提供接口的解密操作
func RSAGet(res []byte) string {
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, res)
	if err != nil {
		logrus.Error("rsa解密内容出错---->", err)
		return ""
	}
	return string(decryptedBytes)
}

// RSAAdd 提供内容的加密操作
func RSAAdd(res string) []byte {
	plain := []byte(res)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, PublicKey, plain)
	if err != nil {
		logrus.Error("rsa加密内容出错---->", err)
	}
	return cipherText
}
