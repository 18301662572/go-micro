package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"go-micro-pro/HTTP-API/Services"
	"go-micro-pro/HTTP-API/Weblib"
	"go-micro-pro/HTTP-API/Wrappers"
)

//http client (httpServer)访问rpc服务(go-micro-rpc)
//1.访问：使用postman 访问client服务
//2.加入go-micro的装饰器wrapper -> logWrapper（中间件），在执行之前做一些其他事情

type logWrapper struct {
	client.Client //导入 "github.com/micro/go-micro/client"
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口")
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	//初始化consul
	consulRegistry := consul.NewRegistry(
		registry.Addrs("http://localhost:8500/"), //consul ip
	)

	//http client服务注册(通过micro)
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper),            //装饰器/拦截器，写日志等...
		micro.WrapClient(Wrappers.NewProdsWrapper), //熔断代码,(处理异常，超时)提供服务的熔断、降级、隔离等。
	)

	//访问服务（grpc service服务名，http client服务名）
	prodService := Services.NewProdService("prodservrice", myService.Client())

	//httpServer 初始化
	httpServer := web.NewService(
		web.Name("httpprodservice"),                   //注册服务名
		web.Address(":8001"),                          //注册端口号   //使用命令行参数启动商品服务,参数是server_address，所以要把这行注释
		web.Handler(Weblib.NewGinRouter(prodService)), //注册Handler
		web.Registry(consulRegistry),                  //注册consul
	)

	httpServer.Init()
	httpServer.Run()
}
