package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 继承 StandardClaims 的struct
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 使用结构体进行秘钥生成
func JWTStruct(username string, mySigningString string) string {

	var myCustomClaims = MyCustomClaims{
		username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    username,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myCustomClaims)
	fmt.Println(t)
	s, err := t.SignedString(mySigningString)
	if err != nil {
		panic(err)
	}
	fmt.Println(s, ", len=", len(s))
	return s
}
