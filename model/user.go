package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int
	Username string `json:"username" validate:"min=4,max=10"`
	Password string `json:"password" validate:"min=6,max=10"`
	jwt.StandardClaims
	//ExpiresAt 和 Issuer 分别为过期时间 和 签发者
	//StandardClaims 嵌入到User 中，以方便对标准声明进行编码、解析和验证
}
asdf