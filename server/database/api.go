package database

import (
	"VocabularyLife/expend/HttpResult"
	"VocabularyLife/expend/cryptoapi"
	"VocabularyLife/expend/jwt"
	"VocabularyLife/server/cacheRedis"
	"VocabularyLife/server/database/flow"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	account, _ := flow.SelectAccount(sign.User, sign.PassWord)

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

// PostAccountEnrollment 注册用户
func PostAccountEnrollment(ctx *gin.Context) {
	// 读取中间件保存在上下文的数据
	sign := ctx.MustGet("sign").(cryptoapi.Signs)
	data := ctx.MustGet("data").(cryptoapi.Datas)
	time := ctx.MustGet("time").(cryptoapi.Times)

	var enrollment Enrollment
	if err := ctx.BindJSON(&enrollment); err != nil {
		logrus.Error("注册账号接口绑定元素错误--->", err)
		ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("注册账号失败，账号密码错误"))
		return
	}
	// 自动生成uid
	uid, err := uuid.NewUUID()
	Uid := fmt.Sprintf("%s", uid)
	if err != nil {
		logrus.Info("生成uuid失败--->", err)
		ctx.JSONP(http.StatusOK, HttpResult.InternalErrorFun("注册账号失败，请刷新"))
	}
	logrus.Printf("请求传入的结果值Sign:%v,Data:%v,Time:%v,Uid:%v", sign, data, time, Uid)

	account, _ := flow.SelectAccount(sign.User, sign.PassWord)

	logrus.Info("请求的账号结构体值->", account)
	// 倘若没有就会返回空
	if account.ID == 0 {
		// 注册账号
		flow.WitheAdminAccount(Uid, sign.User, enrollment.UserName, sign.PassWord, enrollment.Imgurl)
		logrus.Info("注册账号成功：", sign)
		ctx.JSON(http.StatusOK, HttpResult.Success(enrollment, "注册账号成功，请登录账号"))
		return
	}

	logrus.Error("注册账号失败，已有账号:", sign)

	ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("注册账号失败，已有账号"))

}
