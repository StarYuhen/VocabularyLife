package jwt

import (
	"VocabularyLife/configs"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

var config = configs.Config.CryptoConfig

// JwtPO 自定义jwt参数
type jwtPO struct {
	User string `json:"user"`
	Uid  string `json:"uid"`
	jwt.StandardClaims
}

// TokenExpireDuration 设置为100天，一次用户登录有效期
const TokenExpireDuration = time.Hour * 2400

var MySecret = []byte(config.Jwt)

// GenToken 生成JWT
func GenToken(user string, uid string) (string, error) {
	// 创建一个我们自己的声明
	lie := jwtPO{
		user,
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    config.JwtName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, lie)
	tokens, _ := token.SignedString(MySecret)
	logrus.Info("加密过后的jwt---->", tokens)
	return tokens, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwtPO, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &jwtPO{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 将解析后的内容返回
	if clime, ok := token.Claims.(*jwtPO); ok && token.Valid {
		logrus.Info("解析后的jwt内容---->", clime)
		return clime, nil
	}

	return nil, errors.New("invalid token")
}

// AuthMiddlewareJWT  基于JWT的认证中间件
func AuthMiddlewareJWT() func(c *gin.Context) {
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
		// 将当前请求的user和uid信息保存到请求的上下文c上
		c.Set("user", mc.User)
		c.Set("uid", mc.Uid)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
