package router

import (
	"my-go-blog/server/api/sys"
	"my-go-blog/server/middleware"

	"github.com/gin-gonic/gin"
)

func RouterGroupComment(rg *gin.RouterGroup) {
	// 评论路由组
	rgComment := rg.Group("comment").Use(middleware.JWTAuth())
	{
		rgComment.GET("/list/:id", sys.GetCommentsByPostID) // 评论查看
		rgComment.PUT("/create/:id", sys.CreateComment)     // 评论创建
	}
}
