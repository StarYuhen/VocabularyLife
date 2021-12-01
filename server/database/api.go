package database

import (
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/expend/jwt"
	"VocabularyLife/server/cacheRedis"
	"VocabularyLife/server/database/flow"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// POSTAccountLogin  获取当前用户的账号信息，也是登录接口
func POSTAccountLogin(ctx *gin.Context) {

	// 读取中间件保存在上下文的数据
	sign := ctx.MustGet("sign").(cryptoapi.Signs)
	data := ctx.MustGet("data").(cryptoapi.Datas)
	time := ctx.MustGet("time").(cryptoapi.Times)

	logrus.Printf("请求传入的结果值Sign:%v,Data:%v,Time:%v", sign, data, time)
	// logrus.Info("使用反射获取所有类型所有类型，", reflect.TypeOf(sign), reflect.TypeOf(data), reflect.TypeOf(time))

	// 数据库判断是否有该用户 --当查询的值存在时说明有该用户
	account := flow.SelectAccount(sign.User, sign.PassWord)

	if account.User == "" {
		logrus.Error("该用户不存在--->", account)
		ctx.JSON(http.StatusOK, HttpResult.NotOwnedFun("请求的用户不存在"))
		return
	}

	if account.User == sign.User {
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
			flow.WriteUserInfo(account.Uid, data.Ip, jwtGet)
			logrus.Info("插入登录信息数据表")
		} else if user.Uid == account.Uid {
			// 说明不是第一次缓存数据了，需要刷新数据
			flow.UpdateUserInfoIp(account.Uid, data.Ip, jwtGet)
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
