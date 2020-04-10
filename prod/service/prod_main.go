package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"go-micro-pro/helper"
	"go-micro-pro/prod/prodmodel"
)

//商品API服务 （HTTP服务）

//注：使用命令行参数启动商品服务
//go run prod_main.go --server_address :8001   启动一个服务
//go run prod_main.go --server_address :8002   启动另一个服务
//go run prod_main.go --server_address :8003   启动另一个服务

func main() {

	consulRegistry := consul.NewRegistry(
		registry.Addrs("http://localhost:8500/"), //consul ip
	)

	ginRouter := gin.Default()

	v1Group := ginRouter.Group("/v1")
	{

		v1Group.Handle("GET", "/prods", func(context *gin.Context) {
			context.JSON(200, prodmodel.NewProdList(5))
		})

		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			//绑定参数,获取POST访问链接中form传的参数size的值
			var prodsreq helper.ProdsRequest
			err := context.Bind(&prodsreq)
			if err != nil || prodsreq.Size <= 0 {
				prodsreq = helper.ProdsRequest{Size: 2}
			}
			//返回json
			context.JSON(200,
				gin.H{"data": prodmodel.NewProdList(prodsreq.Size)})
		})
	}

	//service 注册
	server := web.NewService(
		web.Name("prodservice"),      //注册服务名
		web.Address(":8001"),         //注册端口号   //使用命令行参数启动商品服务,参数是server_address，所以要把这行注释
		web.Handler(ginRouter),       //注册Handler
		web.Registry(consulRegistry), //注册consul
	)
	server.Init() //使用命令行参数启动商品服务，添加这行
	server.Run()
}
