package cacheRedis

import (
	"VocabularyLife/configs"
	"github.com/go-redis/redis"
)

// redis配置
var config = configs.Config.RedisConfig

var rdb = redisInit()

// 初始化redis链接
func redisInit() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,           // redis数据库index
		PoolSize:     config.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: config.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}
