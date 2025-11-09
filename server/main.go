package main

import (
	"fmt"
	"my-go-blog/server/global"
	sinit "my-go-blog/server/init"
	"my-go-blog/server/model/comm"
	"my-go-blog/server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	sinit.SetDB()
	fmt.Println(global.GVA_DB)
	global.GVA_DB.AutoMigrate(&comm.User{}, &comm.Post{}, &comm.Comment{})

	global.USER_TOKENS = make(map[uint]string)

	r := gin.Default()
	router.RouterBase(r, "goblog")
	// router.RouterGroupUser(rgBase)

	r.Run(":8888")
}
