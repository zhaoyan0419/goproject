package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var s1 []byte = make([]byte, 1024)
		// 从conn连接中读取数据
		n, err := conn.Read(s1)
		if err != nil {

			fmt.Println(conn.RemoteAddr().String(), "连接断开，err：", err)
			return

		}
		//将读取的内容输出在服务器端
		fmt.Println("接收到来自于", conn.RemoteAddr().String(), "的输入", string(s1[:n]))

	}
}

func main() {
	addr := "0.0.0.0:7899"
	// 进行监听，需要指定服务端为tcp协议，监听地址为addr（ip+port）
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("服务端监听失败，err：", err)
		return
	}
	fmt.Println("服务器端监听了:", addr)
	for {
		// 监听成功以后，等待客户端进行连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("等待客户端的连接失败，err：", err)
		} else {
			// 连接成功
			fmt.Printf("等待连接成功，con:%v,接收到的客户端信息: %v\n", conn, conn.RemoteAddr().String())
		}
		// 准备一个协程，协程处理客户端服务请求
		go process(conn) // 不同的客户端的请求，连接conn不一样的
	}

}
