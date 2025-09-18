package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
	"zhaoyan/src/gocode/order/protobufs/compilesReadOnly"
)

var (
	grpcaddr = flag.String("grpc", "localhost:5051", "the address of grpc server")
	addr     = flag.String("addr", "127.0.0.1", "the address for listen .default 127.0.0.1")
	port     = flag.Int("port", 8080, "the port for listen.default is 8080")
)

func main() {
	flag.Parse()
	service := http.NewServeMux()
	service.HandleFunc("/order", func(writer http.ResponseWriter, request *http.Request) {
		// 完成作为grpc客户端的请求
		// 远程调用
		// 连接到grpc服务端:5051
		conn, err := grpc.Dial(*grpcaddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// 实例化grpc客户端
		client := compilesReadOnly.NewProductClient(conn)
		// 远程调用
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		productResp, err := client.ProductInfo(ctx, &compilesReadOnly.ProductRequest{Id: 42})
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Fprintf(writer, resp.Name, resp.Id, resp.IsSale)
		// 构建http的响应
		data := struct {
			ID       int64                               `json:"id"`
			Quantity int                                 `json:"quantity"`
			Products []*compilesReadOnly.ProductResponse `json:"products"`
		}{
			9527, 1,
			[]*compilesReadOnly.ProductResponse{productResp}}

		dataJson, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		writer.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprintf(writer, string(dataJson)); err != nil {
			log.Fatal(err)
		}
	})

	// 启动http监听
	address := fmt.Sprintf("%s:%d", *addr, *port)
	fmt.Println("Order http service is listening on ", address)
	log.Fatal(http.ListenAndServe(address, service))
}
