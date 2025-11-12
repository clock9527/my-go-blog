package router

import (
	"my-go-blog/server/api/sys"
	"my-go-blog/server/middleware"

	"github.com/gin-gonic/gin"
)

func RouterGroupComment(rg *gin.RouterGroup) {
	// 评论路由组
	rgComment := rg.Group("comment")
	{
		rgComment.GET("/getComments/:id", sys.GetCommentsByPostID)                      // 评论查看
		rgComment.PUT("/createCmment/:id", sys.CreateComment).Use(middleware.JWTAuth()) // 评论修改
	}
}
