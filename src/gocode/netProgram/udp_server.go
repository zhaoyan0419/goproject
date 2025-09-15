package netProgram

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
)

const udp = "udp"

// UDPServer监听通用版
func UdpServerBasic() {
	// 解析地址
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5678")
	if err != nil {
		log.Fatal(err)
	}

	// 监听
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("UDP Listening Success, UDP addr is: ", udpConn.LocalAddr().String())
	defer udpConn.Close()
	log.Println(udpConn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		rn, raddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(udpConn.RemoteAddr())
		log.Printf("Read From UDP,The Remote Addr: %s,The Data received: %s\n", raddr.String(), string(buf[:rn]))

		// 写操作

		data := []byte("received:" + string(buf[:rn]))
		wn, err := udpConn.WriteToUDP(data, raddr)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Send %s(%d) to %s\n", string(data), wn, raddr.String())
		log.Println(udpConn.RemoteAddr())
	}

}

// UDP对等连接
func UdpServerPeer() {
	// 解析地址
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5678")
	if err != nil {
		log.Fatal(err)
	}

	// 监听
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("UDP Listening Success, UDP addr is: ", udpConn.LocalAddr().String())
	defer udpConn.Close()

	buf := make([]byte, 1024)
	for {
		rn, raddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Read From UDP,The Remote Addr: %s,The Data received: %s\n", raddr.String(), string(buf[:rn]))

		// 写操作
		data := []byte("received:" + string(buf[:rn]))
		wn, err := udpConn.WriteToUDP(data, raddr)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Send %s(%d) to %s\n", string(data), wn, raddr.String())
	}

}

// 组播接收端
func UDPReceiverMulticast() {
	// 1组播监听地址
	address := "224.1.1.2:6789"
	gaddr, err := net.ResolveUDPAddr(udp, address)
	if err != nil {
		log.Fatal(err)
	}
	// 2 组播监听
	UdpConn, err := net.ListenMulticastUDP(udp, nil, gaddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("接收端监听了组播地址：", gaddr.String())
	// 3 接受数据
	buf := make([]byte, 1024)
	for {

		rn, raddr, err := UdpConn.ReadFromUDP(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println(err)
				break
			}
			log.Println(err)
			continue
		}
		log.Printf("received \"%s\" from %s", string(buf[:rn]), raddr.String())
	}

}

// 广播接收端
func UDPReceiverBroadcast() {
	laddr, err := net.ResolveUDPAddr(udp, ":9999")
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("广播地址已监听：", udpConn.LocalAddr().String())
	defer udpConn.Close()
	buf := make([]byte, 1024)
	for {
		rn, raddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println(err)
				break
			}
			log.Println(err)
			continue
		}
		log.Println("received data: ", string(buf[:rn]), ". From: ", raddr.String())
	}
}

// 文件传输（接收端）
func UDPFileDownLoad() {
	laddress := ":5678"
	laddr, err := net.ResolveUDPAddr(udp, laddress)
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server was listening on", udpConn.LocalAddr().String())
	defer udpConn.Close()

	// 确认文件名并恢复确认
	buf := make([]byte, 4096)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
	}
	filename := string(buf[:rn])
	if _, err := udpConn.WriteToUDP([]byte("filename ok"), raddr); err != nil {
		log.Fatal(err)
	}
	// 确认完毕，创建文件，接收内容
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	i := 1
	for {
		rn, _, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			break
		}
		if string(buf[:rn]) == "EOF" {
			log.Println("文件下载完毕")
			break
		}

		if _, err := file.Write(buf[:rn]); err != nil {
			log.Fatal(err)
		}
		log.Println("read", i)
		i++
	}
}
