package netProgram

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 服务端
func TcpServer() {
	// A. 基于某个地址建立监听
	addr := "127.0.0.1:5678"
	listener, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("监听成功，监听的ip和port是：", listener.Addr())
	defer listener.Close()
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())
	// B. 接收连接请求
	// 因为客户端请求很多，应该是一个循环接受请求
	for {
		// listenr.acccept在请求过来之前是处于阻塞状态
		conn, err := listener.Accept()
		if err != nil {
			// 这里不使用fatalln来打印err，因为当前循环就算错误了也不能将程序停止，不能影响循环主体
			fmt.Println(err)
		}
		// 处理连接
		// 日志连接的远程地址
		log.Printf("accept from %s\n", conn.RemoteAddr().String())
	}

}

// BackLog服务端
func TcpBackLogServer() {
	// A. 基于某个地址建立监听
	addr := "127.0.0.1:5678"
	listener, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("监听成功，监听的ip和port是：", listener.Addr())
	defer listener.Close()
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())
	// B. 接收连接请求
	// 因为客户端请求很多，应该是一个循环接受请求
	for {
		// listenr.acccept在请求过来之前是处于阻塞状态
		conn, err := listener.Accept()
		if err != nil {
			// 这里不使用fatalln来打印err，因为当前循环就算错误了也不能将程序停止，不能影响循环主体
			fmt.Println(err)
		}

		// 使用协程，
		go func(conn net.Conn) {
			// 处理连接
			// 日志连接的远程地址
			log.Printf("accept from %s\n", conn.RemoteAddr().String())
			time.Sleep(time.Second)
		}(conn)

	}

}

// 基本读写操作
func TcpServerRW() {
	// A. 基于某个地址建立监听
	addr := "127.0.0.1:5678"
	listener, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("监听成功，监听的ip和port是：", listener.Addr())
	defer listener.Close()
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())
	// B. 接收连接请求
	// 因为客户端请求很多，应该是一个循环接受请求
	for {
		// listenr.acccept在请求过来之前是处于阻塞状态
		conn, err := listener.Accept()
		if err != nil {
			// 这里不使用fatalln来打印err，因为当前循环就算错误了也不能将程序停止，不能影响循环主体
			fmt.Println(err)
		}

		// 使用协程，
		go HandleConn(conn)

	}
}

// 处理连接请求
func HandleConn(conn net.Conn) {
	log.Printf("accept from %s\n", conn.RemoteAddr().String())
	// 保证连接关闭
	defer conn.Close()
	// 向客户端发送数据
	n, err := conn.Write([]byte("send some data from server" + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("server write len is %d\n", n)
	// 从客户端接收数据
	ServerRead(conn)

}
func ServerRead(conn net.Conn) {
	for {
		s1 := make([]byte, 1024)
		n, err := conn.Read(s1)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("received from client data is :%s", string(s1[:n]))
	}
}
