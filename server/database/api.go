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
	var w flow.WriteIO
	// 读取中间件保存在上下文的数据
	sign := ctx.MustGet("sign").(cryptoapi.Signs)
	data := ctx.MustGet("data").(cryptoapi.Datas)
	time := ctx.MustGet("time").(cryptoapi.Times)

	logrus.Printf("请求传入的结果值Sign:%v,Data:%v,Time:%v", sign, data, time)
	// logrus.Info("使用反射获取所有类型所有类型，", reflect.TypeOf(sign), reflect.TypeOf(data), reflect.TypeOf(time))

	// 数据库判断是否有该用户 --当查询的值存在时说明有该用户

	w.SelectAccount(sign.User, sign.PassWord)

	if w.Admin.User == "" {
		logrus.Error("该用户不存在--->", w.Admin)
		ctx.JSON(http.StatusOK, HttpResult.NotOwnedFun("请求的用户不存在"))
		return
	}

	if w.Admin.User == sign.User {
		// 生成jwt
		jwtGet, err := jwt.GenToken(w.Admin.User, w.Admin.Uid)
		if err != nil {
			logrus.Error("生成jwt失败--->", err)
			ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("生成jwt失败"))
			return
		}

		logrus.Info("登录成功---->", w.Admin)

		// 将内容提交给登录信息数据表

		w.ReadUserInfo()

		// 第一次登录表缓存表数据
		if w.UserInfo.Uid == "" {
			flow.WriteUserInfo(w.UserInfo.Uid, data.Ip, jwtGet)
			logrus.Info("插入登录信息数据表")
		} else {
			// 说明不是第一次缓存数据了，需要刷新数据
			w.UpdateUserInfoIp(data.Ip, jwtGet)
			logrus.Info("更新登录信息数据表")
		}

		// 该用户hash表对应的jwt值
		logrus.Info("修改登录用户的jwt情况:", cacheRedis.HaseSet(w.Admin.Uid, "jwt", jwtGet))

		ctx.JSON(http.StatusOK, HttpResult.Success(w.Admin, jwtGet))
		return
	}

	ctx.JSON(http.StatusOK, HttpResult.UnknownErrorFun("请求参数出现问题，引发未收录错误"))
	return
}

// PostAccountEnrollment 注册用户
func PostAccountEnrollment(ctx *gin.Context) {
	var w flow.WriteIO
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

	w.SelectAccount(sign.User, sign.PassWord)

	logrus.Info("请求的账号结构体值->", w.Admin)
	// 倘若没有就会返回空
	if w.Admin.ID == 0 {
		// 注册账号
		flow.WitheAdminAccount(Uid, sign.User, enrollment.UserName, sign.PassWord, enrollment.Imgurl)
		logrus.Info("注册账号成功：", sign)
		ctx.JSON(http.StatusOK, HttpResult.Success(enrollment, "注册账号成功，请登录账号"))
		return
	}

	logrus.Error("注册账号失败，已有账号:", sign)

	ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("注册账号失败，已有账号"))

}

// PostAccountUpdateImgurl 更新头像
func PostAccountUpdateImgurl(ctx *gin.Context) {
	var w flow.WriteIO
	user := ctx.MustGet("user").(string)
	var enrollment Enrollment
	if err := ctx.Bind(&enrollment); err != nil {
		logrus.Error("传输绑定数据失败", err)
		ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("更新头像失败"))
		return
	}
	w.SelectAccountUser(user)
	w.UpdateAdminAccountImgurl(enrollment.Imgurl)
	if w.Admin.ImgUrl != "" && w.Admin.ImgUrl == enrollment.Imgurl {
		// 说明更新成功
		ctx.JSON(http.StatusOK, HttpResult.Success(enrollment, "更新头像成功"))
		return
	}
	ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("更新头像失败"))
}

// PostAccountPassWord 更新主账号密码
func PostAccountPassWord(ctx *gin.Context) {
	var w flow.WriteIO
	user := ctx.MustGet("user").(string)
	var pass PostAccountPassword
	if err := ctx.Bind(&pass); err != nil {
		logrus.Error("传输绑定数据失败", err)
		ctx.JSON(http.StatusOK, HttpResult.ParameterErrorFun("更新头像失败"))
		return
	}

	w.SelectAccountUser(user)
	w.UpdateAdminAccountPassword(pass.Password)

	if w.Admin.Password == pass.Password {
		ctx.JSON(http.StatusOK, HttpResult.Success(pass, "更新密码成功"))
		return
	}

	ctx.JSON(http.StatusOK, HttpResult.InternalErrorFun("更新密码失败"))
}
