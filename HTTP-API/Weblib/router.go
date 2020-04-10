package Weblib

import (
	"github.com/gin-gonic/gin"
	"go-micro-pro/HTTP-API/Services"
)

func NewGinRouter(prodService Services.ProdService) *gin.Engine {
	//初始化gin
	ginRouter := gin.Default()
	//中间件
	ginRouter.Use(InitMiddleware(prodService), ErrorMiddleware())
	v1Group := ginRouter.Group("/v1")
	{
		//http://192.168.1.4:8001/v1/prods
		v1Group.Handle("POST", "/prods", GetProdList)
		//http://192.168.1.4:8001/v1/prods/123
		v1Group.Handle("GET", "/prods/:pid", GetProdDeatil) //uri:"pid"
	}
	return ginRouter
}
