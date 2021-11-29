package main

import (
	"VocabularyLife/api"
	"VocabularyLife/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("当前VocabularyLife的版本是:", configs.Config.AppVersion)
	// 界面运行展示配置

	// 启动gin接口服务
	api.RoutingApi()
}
