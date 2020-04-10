package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//client  selector获取某个服务的信息
//访问服务：ip+port+路由
//http://192.168.1.4:8001/v1/prods （192.168.1.4:8001：consul服务获取到的address）
//注：如果某个服务停掉，consul（etcd）会监听到服务调开，selector getService获取不到就会断开这个服务
//步骤：
//1.获取注册中心 (consul,etcd)
//2.得到服务（通过服务名）
//3.选择服务（selector 轮询，随机）
//4.调用服务（callapi）

//服务调用：原始方式调用API（http api）
//addr:ip+port path:/v1/prods
func callApi(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf), nil
}

//随机获取服务名下的一个服务
func getSimplateService(consulRegistry registry.Registry) {
	getService, err := consulRegistry.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}

	//selector在go-micro的client下
	next := selector.Random(getService) //随机选择一个服务
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)
}

//轮询获取服务名下的所有服务
func getAllService(consulRegistry registry.Registry) {
	for {
		getService, err := consulRegistry.GetService("prodservice")
		if err != nil {
			log.Fatal(err)
		}

		//selector在go-micro的client下
		next := selector.RoundRobin(getService) //RoundRobin 实现内置的轮询算法实现，但不是平滑
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Id, node.Address, node.Metadata)
		time.Sleep(time.Second * 1)
	}
}

//轮询获取服务名下的一个服务，进行服务调用
func getSimplateServiceCallAPI(consulRegistry registry.Registry) {
	//2.得到服务
	getService, err := consulRegistry.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}

	//3.选择服务
	//selector在go-micro的client下
	next := selector.RoundRobin(getService) //RoundRobin 实现内置的轮询算法实现，但不是平滑
	node, _ := next()

	//4.服务调用
	//直接访问获取到的服务，打印
	callResp, err := callApi(fmt.Sprintf("%s", node.Address), "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(callResp)
	time.Sleep(time.Second)
}

//步骤：
//1.获取注册中心 (consul,etcd)
//2.得到服务（通过服务名）
//3.选择服务（selector 轮询，随机）
//4.调用服务（callapi）
func main() {
	//1.获取注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("http://localhost:8500/"), //consul ip
	)
	//随机获取一个服务
	//getSimplateService(consulRegistry)

	//轮询获取服务名下的所有服务
	//getAllService(consulRegistry)

	//轮询获取服务名下的一个服务，进行服务调用
	getSimplateServiceCallAPI(consulRegistry)
}
