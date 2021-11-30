package cacheRedis

import (
	"github.com/sirupsen/logrus"
	"time"
)

// StringSet 设置redis的字符串及过期时间
func StringSet(key, value string) bool {
	// 统一过期时间60s
	err := rdb.Set(key, value, time.Second*60).Err()
	if err != nil {
		logrus.Info("图片插入redis失败，请刷新")
		return false
	}
	return true
}

// StringRead 获取redis字符串的值
func StringRead(key string) (bool, string) {
	err := rdb.Get(key).Val()
	if err == "" {
		logrus.Info("查询redis失败，请刷新")
		return false, ""
	}
	return true, err
}
