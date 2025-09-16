package netProgram

import (
	"log"
	"net"
	"sync"
	"time"
)

// 网络轮询器

// 网络IO的阻塞
func BIONet() {
	addr := "127.0.0.1:5678"
	wg := sync.WaitGroup{}
	// 模拟读
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conn, err := net.Dial(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		buf := make([]byte, 1024)
		// 注意两次时间间隔
		log.Println("start reading", time.Now().Format("03:04:05.000"))

		rn, err := conn.Read(buf)
		log.Println("content:", string(buf[:rn]), time.Now().Format("03:04:05.000"))
	}(&wg)
	// 模拟写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		listener, err := net.Listen(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				log.Println("connected")
				// 阻塞时长
				time.Sleep(time.Second * 3)
				conn.Write([]byte("Blocking I/O"))
			}(conn)
		}
	}(&wg)
	wg.Wait()
}

// channel阻塞
func BIOChan() {
	// 初始化一个chan
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("read start:", time.Now().Format("03:04:05.000"))
		content := <-ch
		log.Println("content", content, time.Now().Format("03:04:05.000"))
	}(&wg)

	// 写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}(&wg)

	wg.Wait()

}

// 网络IO的非阻塞
func NIONet() {
	addr := "127.0.0.1:5678"
	wg := sync.WaitGroup{}
	// 模拟读
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conn, err := net.Dial(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		buf := make([]byte, 1024)
		// 注意两次时间间隔
		log.Println("start reading", time.Now().Format("03:04:05.000"))
		conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		rn, err := conn.Read(buf)
		log.Println("content:", string(buf[:rn]), time.Now().Format("03:04:05.000"))
	}(&wg)
	// 模拟写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		listener, err := net.Listen(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				log.Println("connected")
				// 阻塞时长
				time.Sleep(time.Second * 3)
				conn.Write([]byte("Blocking I/O"))
			}(conn)
		}
	}(&wg)
	wg.Wait()
}

// channel非阻塞
func NIOChan() {
	// 初始化一个chan
	ch := make(chan struct{ id uint }, 1)
	defer close(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("read start:", time.Now().Format("03:04:05.000"))
		content := struct{ id uint }{}
		select {
		case content = <-ch:
		default:
		}
		log.Println("content", content, time.Now().Format("03:04:05.000"))

	}(&wg)

	// 写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		ch <- struct{ id uint }{42}

	}(&wg)

	wg.Wait()

}

// 网络IO的非阻塞（channel+conn）
func NIONetChannel() {
	addr := "127.0.0.1:5678"
	wg := sync.WaitGroup{}
	// 模拟读
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conn, err := net.Dial(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// 注意两次时间间隔
		log.Println("start reading", time.Now().Format("03:04:05.000"))
		chRead := make(chan []byte)
		wgwg := sync.WaitGroup{}
		wgwg.Add(1)
		go func(conn net.Conn) {
			defer wgwg.Done()
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			chRead <- buf[:n]
		}(conn)
		data := []byte{}
		time.Sleep(100 * time.Millisecond)
		select {
		case data = <-chRead:

		default:

		}

		log.Println("content:", string(data), time.Now().Format("03:04:05.000"))
		wgwg.Wait()
	}(&wg)
	// 模拟写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		listener, err := net.Listen(tcp, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(conn net.Conn) {

				defer conn.Close()
				log.Println("connected")
				// 阻塞时长
				//time.Sleep(time.Second * 3)
				conn.Write([]byte("Blocking I/O"))
			}(conn)
		}
	}(&wg)
	wg.Wait()
}
