package sys

import (
	"my-go-blog/server/global"
	"my-go-blog/server/model/comm"
	"net/http"
	"strconv"

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
	// 通过认证 token 获取 user id
	if mapToken, exists := c.Get("mapToken"); exists {
		userID := (*(mapToken.(*jwt.MapClaims)))["id"]
		post.UserID = uint(userID.(float64)) // 这里为什么是float64？
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
		return
	}
	// post 参数
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(http.StatusFound, gin.H{
			"error": "请确认输入参数",
		})
	}
	// 创建文章
	if err := global.GVA_DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})

}

// 修改文章
func UpdatePost(c *gin.Context) {
	var post comm.Post

	if mapToken, exists := c.Get("mapToken"); exists {
		// 获取 post id
		id := c.Param("id")
		postID, _ := strconv.Atoi(id)

		// 通过认证token获取userID
		tUserID := uint(((*(mapToken.(*jwt.MapClaims)))["id"]).(float64)) // 这里为什么是float64？

		// 验证user和post是否匹配
		if b := checkUserPostMatch(tUserID, postID, &post); !b {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无权限修改该文章",
			})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
		return
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

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})

}

// 删除文章
func DeletePost(c *gin.Context) {
	// 通过认证token获取 user id
	var post comm.Post
	if mapToken, exists := c.Get("mapToken"); exists {
		// 获取 post id
		id := c.Param("id")
		postID, _ := strconv.Atoi(id)

		// 通过认证token获取userID
		tUserID := uint(((*(mapToken.(*jwt.MapClaims)))["id"]).(float64)) // 这里为什么是float64？

		// 验证user和post是否匹配
		if b := checkUserPostMatch(tUserID, postID, &post); !b {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无权限删除该文章",
			})
			return
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效认证",
		})
	}
	// 执行删除
	d := global.GVA_DB.Debug().Where("user_id = ? And id = ?", post.UserID, post.ID).Delete(&post)
	if d.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// 验证user和post是否匹配，并装载post
func checkUserPostMatch(uid uint, pid int, post *comm.Post) bool {
	var userID uint
	// 通过post id 获取 userid，用于比较
	global.GVA_DB.Model(&post).Where("id = ?", pid).Select("user_id").First(&userID)
	if uid == userID {
		post.ID = uint(pid)
		post.UserID = uid
		return true
	}
	return false
}
