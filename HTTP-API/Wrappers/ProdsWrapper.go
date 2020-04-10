package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"go-micro-pro/HTTP-API/Services"
	"strconv"
)

//hystrix 拦截器

type ProdsWrapper struct {
	client.Client
}

//初始化
func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}

//初始化商品
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

//商品列表降级方法
func defaultProds(rsp interface{}) {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 2; i++ {
		models = append(models, newProd(20+i, "pname"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	result.Data = models
}

//通用的降级方法  ：熔断器拦截到error（超时/调用服务出错）,修改返回值rsp (默认返回一个静态数据)
func defaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *Services.ProdListResponse:
		defaultProds(t)
	case *Services.ProdDeatilResponse:
		t.Data = newProd(30, "降级商品")
	}
}

//熔断代码：设置响应不能超过3秒的超时处理
func (p *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//起一个cmd名：（服务名.方法名）
	cmdName := req.Service() + "." + req.Endpoint()
	//1：配置config
	// 设置超时时间，设置熔断器的开启
	// 下面设置的意思：
	// a.如果2次请求，有50%的请求进入降级函数，则开启熔断器，
	// b.开启熔断器状态，不会走调用服务，直接进入降级函数。
	// c.等待5秒再次检测熔断器是否开启。
	configA := hystrix.CommandConfig{
		Timeout:                3000, //响应不能超过3秒
		RequestVolumeThreshold: 2,    //默认20。熔断器请求阙值，意思是有20个请求才进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比默认50（50%）
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒（默认5秒）
	}
	//2：配置command
	hystrix.ConfigureCommand(cmdName, configA)
	//3：执行，使用Do方法
	return hystrix.Do(cmdName, func() error {
		return p.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		//降级方法
		//defaultProds(rsp)
		defaultData(rsp)
		return nil
	})
}
