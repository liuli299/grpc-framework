# grpc-framework
最简单的grpc框架
> 初级教程
```
git clone https://github.com/liuli299/grpc-framework.git
cd grpc-framework
go run server/main.go                 //启动grpc服务, 
                                      //如果被墙不能下载依赖，请使用本地vendor： go run -mod=vendor server/main.go
go run server/proxy.go                //启动代理服务。支持http请求
go run client/main.go                 //grpc客户端请求示例
http://localhost:8080/alie?name=mmd   //http请求示例
```

要求：
1. 使用go mod管理依赖 go version >= 1.11
2. protobuf
3. grpc

目录文件说明：
1. proto：.proto文件目录
2. server：服务端逻辑
  - controller: 具体服务端逻辑实现
  - main.go: grpc服务端入口
  - proxy.go: http请求代理入口
3. client：客户端示例

> 高级教程

如果你想修改proto文件，需要安装一些依赖
1. [protobuf](https://github.com/protocolbuffers/protobuf/releases)
2. [grpc](https://grpc.io/docs/quickstart/go.html)
3. [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

依赖安装好之后，可按照以下步骤实现自己的逻辑
1. 修改.proto文件
2. 执行 sh build.sh
3. 在 server/controller里面添加具体接口实现逻辑
4. 重新运行最开始的命令即可
