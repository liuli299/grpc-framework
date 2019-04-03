# grpc-framework
最简单的grpc框架

> git clone https://github.com/liuli299/grpc-framework.git
> cd grpc-framework
> go run server/main.go                 //启动grpc服务
> go run server/proxy.go                //启动代理服务。支持http请求
> go run client/main.go                 //grpc客户端请求示例
> http://localhost:8080/alie?name=mmd   //http请求示例

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




