package netProgram

import (
	"fmt"
	"log"
	"net"
	"sync"
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

}

// 循环读
func ServerLoopRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
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

// 验证写阻塞
func ServerWriteBlock() {
	ServerAddr := ":5678"
	listener, err := net.Listen(tcp, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务器监听成功，监听了", ServerAddr)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection set false", err)
			break
		}
		go ServerWriteOnce(conn)

	}

}

// 处理conn，循环写入
func ServerLoopWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	// 设置超时，无能无止境的写，永远写
	//conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	num := 10
	data := []byte("hello" + "\n")
	for i := 0; i < num; i++ {

		wn, err := conn.Write(data)
		if err != nil {
			fmt.Println("数据写入失败", err)
			break
		}
		if err == nil && wn == len(data) {
			log.Println("数据写入成功，第", i, "次写入")
		}
	}
}

// 处理conn，单次写入
func ServerWriteOnce(conn net.Conn) {
	data := []byte("server send some data to client" + "\n")
	wn, err := conn.Write(data)
	if err != nil {
		fmt.Println("write failed", err)
		return
	}
	if err == nil && wn == len(data) {
		log.Println("数据写入成功，发送数据长度为", wn)
	}
}

// server监听（通用）
func ServerListen() {
	wg := sync.WaitGroup{}
	ServerAddr := ":5678"
	listener, err := net.Listen(tcp, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Server Listening Success, Listen Address:", ServerAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		wg.Add(2)
		go ServerLoopWrite(conn, &wg)
		go ServerLoopRead(conn, &wg)
	}
}
