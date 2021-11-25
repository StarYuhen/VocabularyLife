package configs

import (
	_ "embed"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configFile []byte

var Config = WriteConfig()

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
