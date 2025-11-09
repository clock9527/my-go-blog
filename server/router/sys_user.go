package router

import (
	"my-go-blog/server/api/sys"

	"github.com/gin-gonic/gin"
)

func RouterGroupUser(router *gin.RouterGroup) {
	rgUser := router.Group("user")
	{
		rgUser.POST("login", sys.Login)  // 登录
		rgUser.POST("reg", sys.Register) // 注册 {"username":"user1","password":"W123456","email":"123@qq.com"}
	}
}
