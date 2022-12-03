package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"message-board/model"
	"time"
)

const TokenExpireDuration = time.Hour * 24

func GenRegisteredClaims(username string) (string, error) {
	// 创建 Claims
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "huihuitui",
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		},
	}
	// 签发人

	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(model.MyKey)
}
func ParesToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	mc := new(model.MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return model.MyKey, nil
	})
	if err != nil { // 解析token失败
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
