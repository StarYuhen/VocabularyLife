package cacheRedis

import "github.com/sirupsen/logrus"

// redis操作hash表内容

// HaseSet 向hash表插入数据 写入已有key会自动更新已有键的值
func HaseSet(key string, findKey string, value string) bool {
	// 向hash数据库读取数据
	res, err := rdb.HSet(key, findKey, value).Result()
	if err != nil {
		logrus.Error("插入redis哈希表失败--->", err)
		return false
	}
	logrus.Info("修改redis的hash表内容是否错误：", res)
	return true
}

// HaseRead 向指定hash表获取对应key的value
func HaseRead(key string, findKey string) bool {
	// 向hash数据库读取数据
	res, err := rdb.HMGet(key, findKey).Result()
	if err != nil {
		logrus.Error("插入redis哈希表失败--->", err)
		return false
	}
	logrus.Info("插入用户数据到redis哈希表内容是否错误：", res)
	return true
}
