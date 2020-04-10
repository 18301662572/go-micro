package Weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-pro/HTTP-API/Services"
	"strconv"
)

//初始化商品
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

//熔断器拦截到error（超时/调用服务出错）,默认返回一个静态数据 defaultProds
func defaultProds() (*Services.ProdListResponse, error) {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 2; i++ {
		models = append(models, newProd(10+i, "pname"+strconv.Itoa(10+int(i))))
	}
	resp := &Services.ProdListResponse{}
	resp.Data = models
	return resp, nil
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

//获取商品详情
func GetProdDeatil(ginCtx *gin.Context) {
	var prodReq Services.ProdRequest
	PanicIfError(ginCtx.BindUri(&prodReq)) //从uri获取参数绑定

	//先走熔断器，如果发生超时，进入降级方法获取静态数据返回给prodResp
	// main.go -> micro.WrapClient(Wrappers.NewProdsWrapper)实现
	prodService := ginCtx.Keys["prodservrice"].(Services.ProdService)
	prodResp, _ := prodService.GetProdDeatil(context.Background(), &prodReq)
	ginCtx.JSON(200, gin.H{
		"data": prodResp.Data,
	})

}

//获取商品列表
func GetProdList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservrice"].(Services.ProdService)
	var prodReq Services.ProdRequest
	err := ginCtx.Bind(&prodReq) //从form获取参数绑定
	if err != nil {
		ginCtx.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {

		//先走熔断器，如果发生超时，进入降级方法获取静态数据返回给prodResp
		// main.go -> micro.WrapClient(Wrappers.NewProdsWrapper)实现
		prodResp, _ := prodService.GetProdList(context.Background(), &prodReq)
		ginCtx.JSON(200, gin.H{
			"data": prodResp.Data,
		})

		////熔断代码改造
		////第一步：配置config
		//configA := hystrix.CommandConfig{
		//	Timeout: 1000,
		//}
		////第二步：配置command
		//hystrix.ConfigureCommand("getprods", configA)
		////第三步：执行，使用Do方法
		//var prodResp *Services.ProdListResponse
		//err := hystrix.Do("getprods", func() error {
		//	//业务代码，调用grpc服务
		//	prodResp, err = prodService.GetProdList(context.Background(), &prodReq)
		//	return err
		//}, func(e error) error {
		//	//降级方法
		//	prodResp, err = defaultProds()
		//	return err
		//})
		//
		//if err != nil {
		//	ginCtx.JSON(500, gin.H{
		//		"status": err.Error(),
		//	})
		//} else {
		//	ginCtx.JSON(200, gin.H{
		//		"data": prodResp.Data,
		//	})
		//}
	}
}
