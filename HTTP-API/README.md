# HTTP-API  go-micro client端
#### 访问go-micoro-rpc 项目下的grpc服务，使用gin框架并对其进行封装。
#### 实现对外REST，对内RPC


```text
1.装饰器wrapper
    go-micro的装饰器wrapper（中间件/拦截器）,目的：在执行目标函数之前做一些事情，比如写日志等...
    https://github.com/micro/go-plugins/tree/master/wrapper
2.熔断器：服务容错（异常，超时处理） 
    hystrix：提供服务的熔断、降级、隔离等。
    https://github.com/afex/hystrix-go 
```


#### 架构
```text
HTTP-REST
    Services
        protos          .proto及生成.pb.go/.micro.go
    Weblib              http client端
        handlers        gin框架   (熔断器处理调用服务超时，在降级方法里添加默认处理)
        middlerwares    gin中间件
        router          gin初始化
    main.go             http client端注册及调用 （装饰器wrapper，调用服务之前写入日志）
```

#### 熔断器的图片
```text
images/wrapper.png
```

#### gen.bat 文件
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