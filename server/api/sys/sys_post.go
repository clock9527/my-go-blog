package sys

import (
	"my-go-blog/server/global"
	"my-go-blog/server/model/comm"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取文章详细信息
func GetDetail(c *gin.Context) {
	var post comm.Post
	postID := c.Param("id")
	if err := global.GVA_DB.First(&post, "id = ?", postID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章查询失败"})
	}
	c.JSON(http.StatusOK, gin.H{
		"post-detail": post,
	})
}

// 获取文章列表
func GetList(c *gin.Context) {
	var posts []comm.Post
	if err := global.GVA_DB.Select("id", "title").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// 创建文章
func CreatePost(c *gin.Context) {

	var post comm.Post

	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(http.StatusFound, gin.H{
			"error": "请确认输入参数",
		})
	}

	// 通过认证 token 获取 user id
	if mapToken, exists := c.Get("mapToken"); exists {
		userID := (*(mapToken.(*jwt.MapClaims)))["id"]
		post.UserID = uint(userID.(float64)) // 这里为什么是float64？
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
	}

	if err := global.GVA_DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
	}

}

// 修改文章
func UpdatePost(c *gin.Context) {
	// 通过认证token获取用户ID
	var post comm.Post
	if mapToken, exists := c.Get("mapToken"); exists {
		userID := (*(mapToken.(*jwt.MapClaims)))["id"]
		post.UserID = uint(userID.(float64)) // 这里为什么是float64？
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
	}
	// 获取文章参数
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无效参数",
		})
		return
	}
	// 执行Update
	if _, err := gorm.G[comm.Post](global.GVA_DB.Debug()).Where("id = ? And user_id = ?", post.ID, post.UserID).Update(c, "content", post.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新失败",
		})
		return
	}

}

// 删除文章
func DeletePost(c *gin.Context) {
	// 通过认证token获取 user id
	var post comm.Post
	if mapToken, exists := c.Get("mapToken"); exists {
		userID := (*(mapToken.(*jwt.MapClaims)))["id"]
		post.UserID = uint(userID.(float64)) // 这里为什么是float64？
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
	}
	// 处理文章ID 并删除
	postID := c.Param("id")
	if err := global.GVA_DB.Debug().Where("user_id = ? And id = ?", post.UserID, postID).Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
