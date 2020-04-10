package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	myhttp "github.com/micro/go-plugins/client/http"
	"log"
)

//HTTP API 访问go-micro client端
//使用go-plugins里面的http调用包
//注：go-plugins/client/http调用包只能通过POST方式访问

func main() {
	//指定Registry
	//etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	consulReg := consul.NewRegistry(
		registry.Addrs("http://localhost:8500/"), //consul ip
	)
	//设置注册中心
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),              //etcdReg
		selector.SetStrategy(selector.RoundRobin), //轮询方式
	)
	getClient := myhttp.NewClient(
		client.Selector(mySelector),
		client.ContentType("application/json"),
	)

	//1.创建request
	req := getClient.NewRequest("httpprodservice", "/v1/prods", map[string]int{"size": 5})
	//2.创建response
	var rsp map[string]interface{}
	err := getClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)

}
