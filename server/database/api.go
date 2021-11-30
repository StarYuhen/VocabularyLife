package database

import (
	"VocabularyLife/configs"
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/expend/jwt"
	"VocabularyLife/server/cacheRedis"
	"VocabularyLife/server/database/flow"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 配置内容信息
var cryptoconfig = configs.Config.CryptoConfig

// POSTAccountLogin  获取当前用户的账号信息，也是登录接口
/*
 请求参数列表：
sign:加密参数 账号+密码，加密前格式： 账号|密码|time时间戳 {user:账号,password:密码,time:3543645654654} RSA
data:当前设备信息 加密前格式: ip信息|1314|time时间戳 {ip:ip地址,msg:1314，time:3543645654654} RSA
time:时间戳加上（520），加密前格式：{time:当前时间戳，msg:加密额外参数=520} RSA
*/
func POSTAccountLogin(ctx *gin.Context) {
	var login Login
	var times Times
	var datas Datas
	var signs Signs
	if err := ctx.BindJSON(&login); err != nil {
		logrus.Error("传参解析结构体失败---->", err)
		ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("发送请求的内容错误"))
		return
	}

	// 将请求的内容进行解密

	// 解析time内容
	timesrc := cryptoapi.RSAGet(login.Time)
	data := cryptoapi.RSAGet(login.Data)
	sign := cryptoapi.RSAGet(login.Sign)

	logrus.Info("请求传入的结果值", string(timesrc), string(data), string(sign))

	// 判断解析传参内容
	if timesrc == nil || data == nil || sign == nil {
		logrus.Error("解析登录加密参数错误")
		ctx.JSON(http.StatusOK, HttpResult.MissingParametersFun("解析登录加密参数错误"))
		return
	}

	// 将内容赋值给结构体
	errTime := json.Unmarshal(timesrc, &times)
	errSign := json.Unmarshal(sign, &signs)
	errData := json.Unmarshal(data, &datas)

	if errData != nil || errSign != nil || errTime != nil {
		logrus.Error("登录接口将内容赋值给结构体失败-->", errTime, errSign, errData)
		ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("登录接口内部逻辑错误"))
		return
	}

	// 复杂至极的判断条件
	if times.Time == datas.Time && datas.Time == signs.Time && times.Msg == cryptoconfig.Time && datas.Msg == cryptoconfig.Data {
		// 当这些都满足时说明校验正常

		// 数据库判断是否有该用户 --当查询的值存在时说明有该用户
		account := flow.SelectAccount(signs.User, signs.PassWord)

		if account.User == "" {
			logrus.Error("该用户不存在--->", account)
			ctx.JSON(http.StatusOK, HttpResult.NotOwnedFun("请求的用户不存在"))
			return
		}

		if account.User == signs.User {
			// 生成jwt
			jwtGet, err := jwt.GenToken(account.User, account.Uid)
			if err != nil {
				logrus.Error("生成jwt失败--->", err)
				ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("生成jwt失败"))
				return
			}

			logrus.Info("登录成功---->", account)

			// 将内容提交给登录信息数据表

			user := flow.ReadUserInfo(account.Uid)

			// 第一次登录表缓存表数据
			if user.Uid == "" {
				flow.WriteUserInfo(account.Uid, datas.Ip, jwtGet)
				logrus.Info("插入登录信息数据表")
			} else if user.Uid == account.Uid {
				// 说明不是第一次缓存数据了，需要刷新数据
				flow.UpdateUserInfoIp(account.Uid, datas.Ip, jwtGet)
				logrus.Info("更新登录信息数据表")
			}

			// 该用户hash表对应的jwt值
			logrus.Info("修改登录用户的jwt情况:", cacheRedis.HaseSet(account.Uid, "jwt", jwtGet))

			ctx.JSON(http.StatusOK, HttpResult.Success(account, jwtGet))
			return
		}

		ctx.JSON(http.StatusOK, HttpResult.UnknownErrorFun("请求参数出现问题，引发未收录错误"))
		return

	}

	logrus.Info("具体的值-->", times, datas, signs)
	ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("请求的参数错误，校验失败"))
	return
}
