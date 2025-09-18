package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"zhaoyan/src/gocode/product/protobufs/compilesReadOnly"
)

var (
	port = flag.Int("port", 5051, "The gRPC Server Port")
)

type ProductServer struct {
	compilesReadOnly.UnimplementedProductServer
}

func (ProductServer) ProductInfo(ctx context.Context, pr *compilesReadOnly.ProductRequest) (*compilesReadOnly.ProductResponse, error) {
	// 基于pr.id的查询工作略，假设查询到了如下数据
	pi := compilesReadOnly.ProductResponse{
		Id:     42,
		Name:   "马士兵 Go 云原生",
		IsSale: true,
	}
	return &pi, nil
}

func main() {
	flag.Parse()
	// 设置TCP的监听器
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	// 将product的服务，注册到grpc服务器中
	compilesReadOnly.RegisterProductServer(gRPCServer, &ProductServer{})
	log.Printf("gRPC Server is listening on %s\n", listener.Addr())
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
