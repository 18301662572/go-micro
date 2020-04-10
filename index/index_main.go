package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

//主站 api

func main() {

	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(200, gin.H{
			"data": data,
		})
	})

	server := web.NewService(
		web.Address(":8000"),   //注册端口号
		web.Handler(ginRouter), //注册Handler
	)

	server.Run()
}
