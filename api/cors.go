package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// 设置允许跨域请求 中间件
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		orgin := context.Request.Header.Get("Origin")
		if orgin != "" {
			// 接收客户端发送的origin （重要！）
			context.Writer.Header().Set("Access-Control-Allow-Origin", orgin)
			// 服务器支持的所有跨域请求的方法
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			context.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			context.Header("Access-Control-Allow-Credentials", "true")
		}

		// 允许类型校验
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic info is: %v\n", err)
			}
		}()

		context.Next()
	}
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			logrus.Info("请求头中auth为空")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			logrus.Info("请求头中auth格式有误")
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			logrus.Info("无效的Token")
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.User)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// 仅使用一次的中间件内容
