package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

// JwtPO 自定义jwt参数
type jwtPO struct {
	User string `json:"user"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 10000

var MySecret = []byte("StarYuhen")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	lie := jwtPO{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "StarYuhen",
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
