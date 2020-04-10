package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	myhttp "github.com/micro/go-plugins/client/http" //http调用包
	Models "go-micro-pro/models"
	"log"
)

//使用go-plugins里面的http调用包
//注：go-plugins/client/http调用包只能通过POST方式访问

//3.调用服务
func callApi2(s selector.Selector) {
	//3.1 创建http client,设置返回内容类型
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	//3.2 封装请求
	//第三个参数是服务端需要接受的参数
	req := myClient.NewRequest("prodservice", "/v1/prods",
		Models.ProdRequest{Size: 5},
	)
	//3.3 创建返回参数
	var rsp Models.ProListResponse
	//3.4 请求调用
	if err := myClient.Call(context.Background(), req, &rsp); err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())

}

func main() {
	//1.获取注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("http://localhost:8500/"), //consul ip
	)
	//2.得到服务，选择服务
	MySelector := selector.NewSelector(
		selector.Registry(consulRegistry),
		selector.SetStrategy(selector.RoundRobin),
	)
	//3.调用服务
	callApi2(MySelector)
}
