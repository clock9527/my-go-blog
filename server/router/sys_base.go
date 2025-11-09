package router

import "github.com/gin-gonic/gin"

func RouterBase(r *gin.Engine, basePath string) *gin.RouterGroup {
	rg := r.Group(basePath)
	return rg
}
