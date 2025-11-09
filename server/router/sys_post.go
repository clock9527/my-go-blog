package router

import "github.com/gin-gonic/gin"

func RouterGroupPost(rg *gin.RouterGroup) {
	rgPost := rg.Group("post")
	{
		// 文章查看
		rgPostQuery := rgPost.Group("query")
		{
			rgPostQuery.POST("/detail") // 文章详细信息
			rgPostQuery.POST("/list")   // 文章列表
		}
		// 文章新增、修改
		rgPostModify := rgPost.Group("modify")
		{
			rgPostModify.POST("/create") // 创建
			rgPostModify.POST("/update") // 修改
			rgPostModify.POST("/delete") // 删除
		}
	}
}
