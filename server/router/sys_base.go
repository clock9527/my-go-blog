package router

import (
	"github.com/gin-gonic/gin"
)

func RouterBase(r *gin.Engine, basePath string) {
	rg := r.Group(basePath) // 根路径路由组
	RouterGroupUser(rg)
	RouterGroupPost(rg)
	RouterGroupComment(rg)
}
