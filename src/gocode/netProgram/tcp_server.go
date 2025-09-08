package netProgram

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
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
		ServerHandleConnConcurrency(conn)
	}
}

// 服务端并发读写
func ServerHandleConnConcurrency(conn net.Conn) {
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ServerLoopRead(conn, &wg)
	wg.Add(1)
	go ServerLoopWrite(conn, &wg, "Server Write From Goroutine1")
	wg.Add(1)
	go ServerLoopWrite(conn, &wg, "Server Write From Goroutine2")
	wg.Add(1)
	go ServerLoopWrite(conn, &wg, "Server Write From Goroutine3")
	wg.Wait()
}

// 服务端循环写入
func ServerLoopWrite(conn net.Conn, wg *sync.WaitGroup, dt string) {
	defer wg.Done()
	// 设置超时，无能无止境的写，永远写
	//conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	num := 10
	data := []byte(dt + "\n")
	for i := 0; i < num; i++ {

		wn, err := conn.Write(data)
		if err != nil {
			fmt.Println("数据写入失败", err)
			break
		}
		if err == nil && wn == len(data) {
			log.Println("数据写入成功，第", i, "次写入,写入内容", dt)
		}
	}
}

// 服务端循环读
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

// server监听格式化数据传输
func ServerListenFormat() {
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
			continue
		}
		ServerHandleConnFormat(conn)
	}
}

// 服务端处理格式化消息并传输
func ServerHandleConnFormat(conn net.Conn) {
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ServerWriteFormat(conn, &wg)
	wg.Wait()
}

// 服务端格式化写入
func ServerWriteFormat(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// 向客户端发送数据
		// 数据编码后发送

		// 创建需要传递的数据
		// 自定义的消息结构类型
		type Message struct {
			ID      uint   `json:"id,omitempty"`
			Code    string `json:"code,omitempty"`
			Content string `json:"content,omitempty"`
		}
		message := Message{
			ID:      rand.Uint(),
			Code:    "SERVER-STANDARD",
			Content: "message from server",
		}

		//// 1. JSON，文本编码
		//// 创建编码器
		//encoder := json.NewEncoder(conn)
		//// 使用编码器进行编码
		//// encode 成功后，会写入到conn，已经完成了conn.write()
		//if err := encoder.Encode(message); err != nil {
		//	log.Println(err)
		//	time.Sleep(time.Second * 2)
		//	continue
		//}
		//log.Println("Message Was Send")
		//time.Sleep(time.Second * 2)

		//// 测试二进制内容
		//var buf bytes.Buffer
		//testEncoder := json.NewEncoder(&buf)
		//testEncoder.Encode(message)
		//fmt.Println(buf.String())
		//conn.Write([]byte(buf.String()))
		//// 2. GOB，Binary编码

		encoder := gob.NewEncoder(conn)
		// 使用编码器进行编码
		// encode 成功后，会写入到conn，已经完成了conn.write()
		if err := encoder.Encode(message); err != nil {
			log.Println(err)
			time.Sleep(time.Second * 2)
			continue
		}
		log.Println("Message Was Send")
		time.Sleep(time.Second * 2)
	}
}

// 短连接编程示例
func ServerListenShort() {
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
			continue
		}
		ServerHandleConnShort(conn)
	}
}

// 服务端处理conn（短连接）
func ServerHandleConnShort(conn net.Conn) {
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ServerWriteShort(conn, &wg)
	wg.Wait()
}

// 服务端短连接写入
func ServerWriteShort(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	// 向客户端发送数据
	// 数据编码后发送

	// 创建需要传递的数据

	message := Message{
		ID:      rand.Uint(),
		Code:    "SERVER-STANDARD",
		Content: "message from server",
	}

	//// 2. GOB，Binary编码

	encoder := gob.NewEncoder(conn)
	// 使用编码器进行编码
	// encode 成功后，会写入到conn，已经完成了conn.write()
	if err := encoder.Encode(message); err != nil {
		log.Println(err)
		return
	}
	log.Println("Message Was Send")
	return
}

// HeartBeat心跳检测编程示例
func ServerListenHB() {
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
			continue
		}
		go ServerHandleConnHB(conn)
	}
}

// 服务端处理conn（心跳检测）
func ServerHandleConnHB(conn net.Conn) {
	fmt.Println("Connection is establish with ", conn.RemoteAddr().String())
	defer conn.Close()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go ServerPing(conn, &wg)

	wg.Wait()
}

// 服务端写入（心跳检测）
func ServerPing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithCancel(context.Background())
	go ServerReadPong(conn, ctx)
	const maxPingNum = 3
	pingErrCounter := 0
	// 周期性的发送
	// 利用time.ticker
	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		pingMsg := Message{
			ID:   rand.Uint(),
			Code: "SERVER-PING",
			Time: t,
		}
		//// 2. GOB，Binary编码
		encoder := gob.NewEncoder(conn)
		if err := encoder.Encode(pingMsg); err != nil {

			fmt.Println("第", pingErrCounter+1, "次心跳检测失败，连接即将断开")
			pingErrCounter++
			if pingErrCounter == maxPingNum {
				log.Println("连接断开")
				return
				cancel()
			}
		}
		log.Println("ping send to ", conn.RemoteAddr().String(), "ID", pingMsg.ID)
	}

}

func ServerReadPong(conn net.Conn, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message := Message{}
			decoder := gob.NewDecoder(conn)
			err := decoder.Decode(&message)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println(err)
					break
				}
				log.Println(err)
				continue
			}
			log.Println(message.Code)
		}

	}
}

// ConnPool连接池服务端测试代码
func ServerListenPool() {
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
			continue
		}
		go ServerHandlePool(conn)
	}
}

// 服务端处理conn（心跳检测）
func ServerHandlePool(conn net.Conn) {
	fmt.Println("Connection is establish with ", conn.RemoteAddr().String())
	defer conn.Close()
	select {}
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go ServerPing(conn, &wg)
	//
	//wg.Wait()
}

// 服务端写入（心跳检测）
func ServerPingPool(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithCancel(context.Background())
	go ServerReadPong(conn, ctx)
	const maxPingNum = 3
	pingErrCounter := 0
	// 周期性的发送
	// 利用time.ticker
	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		pingMsg := Message{
			ID:   rand.Uint(),
			Code: "SERVER-PING",
			Time: t,
		}
		//// 2. GOB，Binary编码
		encoder := gob.NewEncoder(conn)
		if err := encoder.Encode(pingMsg); err != nil {

			fmt.Println("第", pingErrCounter+1, "次心跳检测失败，连接即将断开")
			pingErrCounter++
			if pingErrCounter == maxPingNum {
				log.Println("连接断开")
				return
				cancel()
			}
		}
		log.Println("ping send to ", conn.RemoteAddr().String(), "ID", pingMsg.ID)
	}

}

func ServerReadPongPool(conn net.Conn, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message := Message{}
			decoder := gob.NewDecoder(conn)
			err := decoder.Decode(&message)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println(err)
					break
				}
				log.Println(err)
				continue
			}
			log.Println(message.Code)
		}

	}
}

// 粘包测试代码
func TcpServerSticky() {
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
			continue
		}
		go HandleConnSticky(conn)
	}
}

// 粘包测试handleConn
func HandleConnSticky(conn net.Conn) {
	fmt.Println("Connection is establish with ", conn.RemoteAddr().String())
	defer conn.Close()
	for i := 0; i < 10; i++ {
		data := "package data."
		_, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println("Write failed: ", err)
			break
		}
	}
}
