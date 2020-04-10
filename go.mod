module go-micro-pro

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/favadi/protoc-go-inject-tag v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5
	github.com/google/btree v1.0.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/protoc-gen-micro v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200408040146-ea54a3c99b9b // indirect
	golang.org/x/tools v0.0.0-20200408032209-46bd65c8538f // indirect
)

replace github.com/micro/go-micro v1.16.0 => github.com/micro/go-micro v1.13.2 //替换版本号，1.18.0没有consul,selector包
