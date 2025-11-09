package router

import "github.com/gin-gonic/gin"

func RouterGroupComment(rg *gin.RouterGroup) {
	// 评论路由组
	rgComment := rg.Group("comment")
	{
		rgComment.POST("/query")  // 评论查看
		rgComment.POST("/create") // 评论修改
	}
}
