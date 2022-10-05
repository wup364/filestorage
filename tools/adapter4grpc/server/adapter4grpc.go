package main

import (
	"flag"
	"fmt"
	"fstools/adapter4grpc/server/service"
	"fstools/adapter4grpc/server/service/impl"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenAddr := flag.String("listen", ":5059", "listen address, default :5059")
	flag.Parse()

	//
	if len(*listenAddr) == 0 {
		panic("listen address can not empty!")
	}else{
		fmt.Printf("listen: %s \r\n", *listenAddr)
	}

	// 监听本地端口

	if lis, err := net.Listen("tcp", *listenAddr); err != nil {
		panic(fmt.Sprintf("监听端口失败: %s", err))
	} else {
		// 创建gRPC服务器
		s := grpc.NewServer()
		// 注册服务
		service.RegisterOpenApiServer(s, &impl.OpenApiServerImpl{})
		reflection.Register(s)
		//
		if err = s.Serve(lis); err != nil {
			panic(fmt.Sprintf("开启服务失败: %s", err))
		}
	}
}
