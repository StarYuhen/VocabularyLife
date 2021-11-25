package configs

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// 获取配置文件内容
//go:embed config.yaml
var configFile []byte

// 获取rsa加密的私钥
//go:embed private.pem
var private []byte

// 获取rsa加密的公钥
// go:embed public.pem
// var public []byte

var Config = WriteConfig()

// PrivateKey rsa加密密钥--私钥
var PrivateKey = PrivateKeyFun()

// WriteConfig 读取配置文件内容写入结构体
func WriteConfig() config {
	var con config
	// 将文件内容映射进入结构体
	if err := yaml.UnmarshalStrict(configFile, &con); err != nil {
		logrus.Error("读取文件映射结构体出错")
		panic(err)
	}
	logrus.Info("读取到的yaml配置文件的值----->", con)
	return con
}

func PrivateKeyFun() *rsa.PrivateKey {
	block, _ := pem.Decode(private)
	// X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}
