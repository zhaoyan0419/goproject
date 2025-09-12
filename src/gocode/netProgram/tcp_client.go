package netProgram

import (
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

// 客户端
func TcpClient() {
	ServerAddr := "127.0.0.1:5678" // 如果省略端口，会随机使用一个端口
	// A. 建立连接
	// 模拟多客户端
	// 并发的客户端请求
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.Dial(tcp, ServerAddr)
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()
			log.Printf("connection is established ,client addr is %s\n", conn.LocalAddr())
		}(&wg)
	}
	wg.Wait()

}

// TimeOutClient客户端
func TcpTimeOutClient() {
	ServerAddr := "127.0.0.1:22" // 如果省略端口，会随机使用一个端口
	// A. 建立连接
	// 模拟多客户端
	// 并发的客户端请求
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.DialTimeout(tcp, ServerAddr, time.Second)
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()
			log.Printf("connection is established ,client addr is %s\n", conn.LocalAddr())
		}(&wg)

	}
	wg.Wait()

}

// 客户端
func TcpBackLogClient() {
	ServerAddr := "127.0.0.1:5678" // 如果省略端口，会随机使用一个端口
	// A. 建立连接
	// 模拟多客户端
	// 并发的客户端请求
	num := 256
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {

		go func(wg *sync.WaitGroup, no int) {
			defer wg.Done()
			conn, err := net.DialTimeout(tcp, ServerAddr, time.Second)
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()
			log.Printf("%d: connection is established ,client addr is %s\n", no, conn.LocalAddr())
		}(&wg, i)
		time.Sleep(30 * time.Millisecond)
	}
	wg.Wait()

}

func TcpClientRW() {
	ServerAddr := "127.0.0.1:5678" // 如果省略端口，会随机使用一个端口
	// A. 建立连接
	// 模拟多客户端
	// 并发的客户端请求
	num := 2
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {

		go func(wg *sync.WaitGroup, no int) {
			defer wg.Done()
			conn, err := net.DialTimeout(tcp, ServerAddr, time.Second)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("%d: connection is established ,client addr is %s\n", no, conn.LocalAddr())
			ClientRead(conn)

		}(&wg, i)
		time.Sleep(30 * time.Millisecond)
	}
	wg.Wait()

}

// 客户端读取conn中内容
func ClientRead(conn net.Conn) {
	defer conn.Close()
	ClientWrite(conn)
	for {
		s1 := make([]byte, 1024)
		n, err := conn.Read(s1)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("received from server data is :%s\n", string(s1[:n]))

	}

}
func ClientWrite(conn net.Conn) {
	n, err := conn.Write([]byte("send some data to server" + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("client write len is %d\n", n)
}

// 验证写阻塞的client
func ClientNoneRead() {
	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial(tcp, ServerAddr)
	if err != nil {
		log.Fatal("dial false", err)
	}
	defer conn.Close()
	log.Println("dial success,client addr is ", conn.LocalAddr().String())
	ClientLoopRead(conn, &sync.WaitGroup{})
}

// 单次读取
func ClientReadOnce(conn net.Conn) {
	data := make([]byte, 1024)
	rn, err := conn.Read(data)
	if err != nil {
		fmt.Println("read failed", err)
		return
	}
	fmt.Println("read data seccess,data:", string(data[:rn]))
}

// Client拨号dial（通用）
func ClientDial() {

	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial(tcp, ServerAddr)
	if err != nil {
		log.Fatal("Dial failed", err)
	}
	defer conn.Close()
	fmt.Println("Connect Server Success,ServerAddr:", conn.RemoteAddr().String())
	ClientHandleConnConcurrency(conn)

}

// 客户端并发处理conn
func ClientHandleConnConcurrency(conn net.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ClientLoopRead(conn, &wg)
	wg.Add(1)
	go ClientLoopWrite(conn, &wg, "Client Write From Goroutine1")
	wg.Add(1)
	go ClientLoopWrite(conn, &wg, "Client Write From Goroutine2")
	wg.Add(1)
	go ClientLoopWrite(conn, &wg, "Client Write From Goroutine3")
	wg.Wait()
}

// 客户端循环写
func ClientLoopWrite(conn net.Conn, wg *sync.WaitGroup, dt string) {
	defer wg.Done()
	WriteData := []byte(dt + "\n")
	num := 10
	for i := 0; i < num; i++ {
		n, err := conn.Write(WriteData)
		if err != nil {
			fmt.Println("Write failed", err)
		}
		fmt.Println(i, ",Write data lenth:", n, dt)
	}
}

// 客户端循环读取
func ClientLoopRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	data := make([]byte, 1024)
	for {
		rn, err := conn.Read(data)
		if err != nil {
			log.Println("read data error:", err)
			break
		}
		fmt.Println("get some data from server:", string(data[:rn]))
	}
}

// Client拨号dial，格式化数据传输
func ClientDialFormat() {

	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial(tcp, ServerAddr)
	if err != nil {
		log.Fatal("Dial failed", err)
	}
	defer conn.Close()
	fmt.Println("Connect Server Success,ServerAddr:", conn.RemoteAddr().String())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ClientReadFormat(conn, &wg)
	wg.Wait()

}

// 客户端传输格式化消息
func ClientHandleConnFromat(conn net.Conn) {
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ClientReadFormat(conn, &wg)
	wg.Wait()
}

// 客户端格式化读消息
func ClientReadFormat(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// 从客户端接收到数据
		//接收到数据后，解码后使用
		type Message struct {
			ID      uint   `json:"id,omitempty"`
			Code    string `json:"code,omitempty"`
			Content string `json:"content,omitempty"`
		}
		//传递的消息类型
		message := Message{}
		//// 1. JSON解码
		//decoder := json.NewDecoder(conn)
		//// 利用解码器进行解码
		//// 解码操作，从conn中读取内容，成功后会将解码后的结果赋值到message
		//if err := decoder.Decode(&message); err != nil {
		//	log.Println(err)
		//	continue
		//}
		//log.Println(message)

		// 2. GOB解码
		decoder := gob.NewDecoder(conn)
		// 利用解码器进行解码
		// 解码操作，从conn中读取内容，成功后会将解码后的结果赋值到message
		if err := decoder.Decode(&message); err != nil {
			log.Println(err)
			continue
		}
		log.Println(message)
	}
}

// Client拨号dial，短连接编程示例
func ClientDialShort() {
	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial(tcp, ServerAddr)
	if err != nil {
		log.Fatal("Dial failed", err)
	}
	fmt.Println("Connect Server Success,ServerAddr:", conn.RemoteAddr().String())
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ClientReadShort(conn, &wg)
	wg.Wait()
}

// 短连接读
func ClientReadShort(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	for {
		// 从客户端接收到数据
		//接收到数据后，解码后使用
		//传递的消息类型
		message := Message{}

		// 2. GOB解码
		decoder := gob.NewDecoder(conn)
		// 利用解码器进行解码
		// 解码操作，从conn中读取内容，成功后会将解码后的结果赋值到message
		err := decoder.Decode(&message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Link Was Closed")
				break
			}
			log.Println(err)
			continue
		}
		log.Println(message)
	}
}

// 相应服务端长连接
// Client拨号dial，短连接编程示例
func ClientDialHB() {
	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial(tcp, ServerAddr)
	if err != nil {
		log.Fatal("Dial failed", err)
	}
	defer conn.Close()
	fmt.Println("Connect Server Success,ServerAddr:", conn.RemoteAddr().String())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go ClientReadPing(conn, &wg)
	wg.Wait()
}

// 短连接读
func ClientReadPing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	for {
		message := Message{}

		// 2. GOB解码
		decoder := gob.NewDecoder(conn)
		// 利用解码器进行解码
		// 解码操作，从conn中读取内容，成功后会将解码后的结果赋值到message
		err := decoder.Decode(&message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Link Was Closed")
				break
			}
			log.Println(err)
			continue
		}
		// 判断是否为ping消息类型
		if message.Code == "SERVER-PING" {
			log.Println("receive ping from ", conn.RemoteAddr().String())
			ClientWritePong(conn, message)
		}
	}
}

func ClientWritePong(conn net.Conn, msg Message) {
	message := Message{
		ID:      rand.Uint(),
		Code:    "CLIENT-PONG",
		Content: fmt.Sprintf("这个pong是用来回复ID为%d这个ping的", msg.ID),
		Time:    time.Now(),
	}
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("pong was send to ", conn.RemoteAddr().String(), msg.Content)
	return
}

// 使用连接池的客户端代码示例

func ClientUseTcpPool() {
	ServerAddr := "127.0.0.1:5678"
	// 建立连接池
	pool, err := NewTcpPool(ServerAddr, PoolConfig{
		InitConnNum: 1,
		MaxConnNum:  10,
		MaxIdleNum:  1,
		IdleTimeout: time.Second * 10,
		Factory:     &TcpConnFactory{},
	})
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 获取连接
			conn, err := pool.Get()
			if err != nil {
				log.Println(err)
				return
			}
			//log.Println(conn)
			// 回收连接
			pool.Put(conn)
		}(&wg)
	}
	//conn, err := pool.Get()
	//defer pool.Put(conn)
	wg.Wait()
	// 释放连接池
	pool.Release()

}

// 粘包客户端read操作
func TcpClientSticky() {
	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial("tcp", ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		rn, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(buf[:rn]))
	}
}

// 客户端解决read操作粘包
func TcpClientCoder() {
	ServerAddr := "127.0.0.1:5678"
	conn, err := net.Dial("tcp", ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	decoder := NewDecoder(conn)
	var data string
	for {
		err := decoder.Decode(&data)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(data)
	}
}

// Tcp专用方法测试读
func TcpClientSpecial() {
	ServerAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5678")
	if err != nil {
		log.Fatal(err)
	}
	TcpConn, err := net.DialTCP(tcp, nil, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TcpConn Set Success")
	defer TcpConn.Close()

	data := make([]byte, 1024)
	for {
		rn, err := TcpConn.Read(data)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println("Read Failed", err)
			continue
		}
		log.Println("Read Success,data:", string(data[:rn]))
	}

}
