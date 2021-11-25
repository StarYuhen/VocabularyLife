package main

import (
	"VocabularyLife/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	// 界面运行展示配置
	logrus.Info("当前VocabularyLife的版本是:", configs.Config.AppVersion)
}
