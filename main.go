package main

import (
	"VocabularyLife/api"
	"VocabularyLife/configs"
	"github.com/sirupsen/logrus"
)

var config = configs.Config

func main() {
	// 运行展示配置服务打印配置
	logrus.Info("当前VocabularyLife的版本是:", config.AppVersion)
	logrus.Printf("gin启动服务,ip地址:%s,开放端口:%s", config.HTTPConfig.Ip, config.HTTPConfig.Port)
	logrus.Printf("加密RSA参数启动服务,Time字段加密值:%d,Data字段加密值:%d", config.CryptoConfig.Time, config.CryptoConfig.Data)
	logrus.Printf("图形验证码启动服务,数字验证码数字长度:%d,图形长度:%d,图形宽度:%d", config.CaptchaConfig.NumberKeyLong, config.CaptchaConfig.ImgHeight, config.CaptchaConfig.ImgWidth)
	logrus.Printf("MYSQL启动服务,MYSQL基础配置:%s,mysql所有数据表:%s", config.MysqlConfig.UserConfig, config.MysqlConfig.UserDBTable)
	logrus.Printf("Redis启动服务,地址:%s,密码:%s,数据库:%d,连接池:%d,最大空闲连接数:%d", config.RedisConfig.Addr, config.RedisConfig.Password, config.RedisConfig.DB, config.RedisConfig.PoolSize, config.RedisConfig.MinIdleConns)

	// 启动gin接口服务
	api.RoutingApi()
}
