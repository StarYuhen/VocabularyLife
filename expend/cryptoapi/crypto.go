package cryptoapi

import (
	"VocabularyLife/configs"
	"crypto"
	"crypto/rsa"
)

// PrivateKet 恒久不变的私钥内容
var PrivateKet = configs.PrivateKey

// RSAGet  提供接口的解密操作
func RSAGet(res []byte) string {
	decryptedBytes, err := PrivateKet.Decrypt(nil, res, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
}
