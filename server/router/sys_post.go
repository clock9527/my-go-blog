package router

import (
	"my-go-blog/server/api/sys"
	"my-go-blog/server/middleware"

	"github.com/gin-gonic/gin"
)

func RouterGroupPost(rg *gin.RouterGroup) {
	rgPost := rg.Group("post")
	{
		// 文章查看
		rgPostQuery := rgPost.Group("query")
		{
			rgPostQuery.GET("/detail/:id", sys.GetDetail) // 文章详细信息
			rgPostQuery.GET("/list", sys.GetList)         // 文章列表
		}
		// 文章新增、修改
		rgPostModify := rgPost.Group("modify")
		rgPostModify.Use(middleware.JWTAuth()) // 添加JWT认证中间件
		{
			rgPostModify.POST("/create", sys.CreatePost)       // 创建
			rgPostModify.POST("/update", sys.UpdatePost)       // 修改
			rgPostModify.DELETE("/delete/:id", sys.DeletePost) // 删除
		}
	}
}
