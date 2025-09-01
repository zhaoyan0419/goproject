package netProgram

import (
	"log"
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
