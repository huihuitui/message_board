package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	// 可根据需要自行添加字段
	Username string `json:"username"`
	jwt.StandardClaims
}
