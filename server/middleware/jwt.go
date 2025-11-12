package middleware

import (
	"fmt"
	"my-go-blog/server/global"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token,Authorization:{{}}
		tokenSecretStr := c.GetHeader("Authorization")
		if tokenSecretStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证token"})
			c.Abort()
			return
		}

		// 解析token
		token, err := jwt.ParseWithClaims(tokenSecretStr, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return global.JWT_KEY, nil
		})

		// token解析是否异常
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "认证失败"})
			}
			c.Abort()
			return
		}
		// token 是否失效
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token失效"})
			c.Abort()
		}

		fmt.Println("exp:", (*(token.Claims.(*jwt.MapClaims)))["exp"])
		c.Set("mapToken", token.Claims)
		c.Next()
	}
}
