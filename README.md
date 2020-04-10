# go-micro微服务架构

## HTTP-API 项目 :
    http api服务，是go-micro client端，访问server服务:go-micoro-rpc 项目下的rpc服务，跟本项目没有关系
## HTTP REST 项目:
    外部客户端http调用 http api服务，访问go-micro client端，跟本项目没有关系

## go-micro-http api项目
### git仓库:https://github.com/micro
### 解析图
```text
images/http-api.png
```

#### 1.架构
```text
hepler
    prods-helper.go     商品数量参数（post传参）
index
    index_main.go       主站API服务
models
    protos
        Prods.proto     商品proto文件
prod
    client-api
        client.go       Client原始方式调用API服务
    client-http-api 
        client-http.go  Client使用go-plugins里面的http调用包调用API服务
    prodmodel
        prodmodel.go    商品信息
    service 
        prod_main.go    Service商品API服务
go.mod
prod.bat                命令行，开启多个prod_main服务
gen.bat                 proto文件生成go文件的命令
HTTP-API                **client端，访问go-micoro-rpc 项目下的rpc服务，跟本项目没有关系**
HTTP-REST               **http api客户端访问 go-micro client端服务，跟本项目没有关系**
```

#### 2.安装插件
```text
//安装go-micro插件
go get -u github.com/micro/go-micro (使用go mod 标注版本号)

//安装gin插件
go get -u github.com/gin-gonic/gin

//安装可选插件
go get -u github.com/micro/go-plugins

//安装grpc插件
1.下载protobuf及protoc-gen-go工具，exe工具
2.下载插件
go get github.com/micro/protoc-gen-micro


//安装第三方插件,生成protoc-go-inject-tag.exe工具
  处理参数模型中的json tag不一致的问题(或者加vaild验证也可以)
  http://github.com/favadi/protoc-go-inject-tag
  go get -u github.com/favadi/protoc-go-inject-tag        
```

#### 3.gen.bat 文件
```text
gen.bat                 
    proto文件生成go文件的命令
         1.cd 进入文件夹下
         2.protoc使用插件执行.proto文件，将生成的go文件存放在models目录下（../表示上级目录）
         3.运行protoc-go-inject-tag 生成的文件.pb.go,改写.pb.go对应的json字段,form字段
         4.cd .. & cd .. 回到当前项目文件夹下（go-micro-pro）
         5.执行，终端直接执行gen.bat
注：go_out 生成的是 .pb.go文件 （模型结构体文件）                                   --message
    micro_out 生成的是 .micro.go文件(服务接口文件interface和服务注册Handler)        --service
```

