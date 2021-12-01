package main

import (
	"VocabularyLife/api"
	"VocabularyLife/configs"
	"github.com/sirupsen/logrus"
)

var config = configs.Config

func main() {
	// 运行展示配置服务打印配置
	logrus.Printf("当前VocabularyLife的后端版本:%d,前端版本:%d", config.AppVersion.Backend, config.AppVersion.Web)
	logrus.Printf("Gin启动服务,ip地址:%s,开放端口:%s", config.HTTPConfig.Ip, config.HTTPConfig.Port)
	logrus.Printf("加密RSA参数接口启动服务,Time字段加密值:%d,Data字段加密值:%d", config.CryptoConfig.Time, config.CryptoConfig.Data)
	logrus.Printf("图形验证码启动服务,数字验证码数字长度:%d,图形长度:%d,图形宽度:%d", config.CaptchaConfig.NumberKeyLong, config.CaptchaConfig.ImgHeight, config.CaptchaConfig.ImgWidth)
	logrus.Printf("MYSQL启动服务,MYSQL基础配置:%s,MYSQL所有数据表:%s", config.MysqlConfig.UserConfig, config.MysqlConfig.UserDBTable)
	logrus.Printf("Redis启动服务,地址:%s,密码:%s,数据库:%d,连接池:%d,最大空闲连接数:%d", config.RedisConfig.Addr, config.RedisConfig.Password, config.RedisConfig.DB, config.RedisConfig.PoolSize, config.RedisConfig.MinIdleConns)
	logrus.Printf("Jwt的加密参数:%s,Jwt的签署人:%s", config.CryptoConfig.Jwt, config.CryptoConfig.JwtName)
	logrus.Printf("Cookie和Sessions的加密参数:%s,Cookie和Sessions的储存名:%s,Cookie和Sessions的参数名列表%s", config.CryptoConfig.Cookie, config.CryptoConfig.CookieName, config.CryptoConfig.CookieList)
	// 启动gin接口服务
	api.RoutingApi()
}
