package cacheRedis

import (
	"VocabularyLife/configs"
	"github.com/go-redis/redis"
)

// redis配置
var config = configs.Config.RedisConfig

// 链接数据库0
var rdb = redisInit()

// 链接数据库1
// var rdbe = redisInitExpend()

// var ctx = context.Background()

// 初始化redis链接
func redisInit() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB[0],        // redis数据库index
		PoolSize:     config.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: config.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}

/*
func redisInitExpend() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB[1],        // redis数据库index
		PoolSize:     config.PoolSize,     // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: config.MinIdleConns, // 初始规定的redis，维护，让其不少于这个数
	})
}

*/
