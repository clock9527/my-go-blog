package sys

import (
	"my-go-blog/server/global"
	"my-go-blog/server/model/comm"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 基于文章ID获取评论
func GetCommentsByPostID(c *gin.Context) {
	var comments []comm.Comment
	if postID, err := strconv.Atoi(c.Param("id")); err != nil { // 获取文章ID
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文章ID参数无效",
		})
		return
	} else if err = global.GVA_DB.Debug().Find(&comments, "post_id=?", postID).Error; err != nil { //获取评论
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "评论获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": comments,
	})

}

// 创建评论
func CreateComment(c *gin.Context) {
	var comment comm.Comment
	// 通过认证 token 获取 user id
	if mapToken, exists := c.Get("mapToken"); exists {
		userID := (*(mapToken.(*jwt.MapClaims)))["id"]
		comment.UserID = uint(userID.(float64)) // 这里为什么是float64？
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
		return
	}
	// 获取评论内容
	if err := c.ShouldBindBodyWithJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "评论内容参数无效",
		})
		return
	}

	// 获取文章ID
	if postID, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文章ID参数无效",
		})
		return
	} else {
		comment.PostID = uint(postID)
	}

	if err := global.GVA_DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": "评论创建失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "创建成功"})
	}

}
