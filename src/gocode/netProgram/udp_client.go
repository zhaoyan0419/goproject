package netProgram

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// UDP客户端连接基础版
func UdpClientBasic() {
	raddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:5678")
	if err != nil {
		log.Fatal(err)
	}
	UdpConn, err := net.DialUDP(udp, nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Udp Set Success", UdpConn.RemoteAddr())
	defer UdpConn.Close()
	// 写
	data := []byte("client write some data to Server")
	wn, err := UdpConn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(UdpConn.RemoteAddr())
	log.Println("Client Send Some Data To Server,Len:", wn)
	defer UdpConn.Close()
	// 读
	buf := make([]byte, 1024)
	rn, raddr, err := UdpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client Received Data From Server: ", string(buf[:rn]))
	log.Println(UdpConn.RemoteAddr())
}

// UDP对等连接测试
func UdpClientPeer() {
	// 解析地址
	raddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:5678")
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8765")
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

	// 写操作
	data := []byte("客户端发送的对等连接的写操作")
	wn, err := udpConn.WriteToUDP(data, raddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Send %s(%d) to %s\n", string(data), wn, raddr.String())

	// 读操作
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read From UDP,The Remote Addr: %s,The Data received: %s\n", raddr.String(), string(buf[:rn]))

}

// 组播发送端
func UDPSenderMulticast() {
	// 建立多播组连接
	raddr, err := net.ResolveUDPAddr(udp, "224.1.1.2:6789")
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.DialUDP(udp, nil, raddr)
	defer udpConn.Close()
	for {
		data := fmt.Sprintf("[%s]: %s\n", time.Now().Format("03:04:05.000"), "hello!")
		_, err = udpConn.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("向组播地址发送了", string(data))

		time.Sleep(time.Second)
	}

}

// 广播发送端

// 广播发送端
func UDPSenderBroadcast() {
	raddr, err := net.ResolveUDPAddr(udp, "192.158.1.255:9999")
	if err != nil {
		log.Fatal(err)
	}
	laddr, err := net.ResolveUDPAddr(udp, ":8888")
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("发送端已就绪：", udpConn.LocalAddr().String())
	defer udpConn.Close()
	// 写
	i := 1

	for {
		data := []byte("向广播地址发送数据。。。" + strconv.Itoa(i))
		_, err := udpConn.WriteToUDP(data, raddr)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("数据发送成功，第", i, "次发送")
		i++
	}

}

// 文件传输（上传）
func UDPFileUpLoad() {
	// 1 获取文件信息
	filename := "./xxx.mp3"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 文件信息
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	//fileinfo.Size()	fileinfo.Name()

	// 2 连接服务器
	raddress := "192.168.1.14:5678"
	raddr, err := net.ResolveUDPAddr(udp, raddress)
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.DialUDP(udp, nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer udpConn.Close()

	// 3 发送文件名
	if _, err := udpConn.Write([]byte(fileinfo.Name())); err != nil {
		log.Fatal(err)
	}
	// 4 服务器确认
	buf := make([]byte, 4096)
	rn, err := udpConn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	// 判断服务端
	if string(buf[:rn]) != "filename ok" {
		log.Println(errors.New("filename error"))
		return
	}
	i := 1
	// 发送文件内容
	for {
		rn, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("file read over")
				udpConn.Write([]byte("EOF"))
				break
			}
			log.Fatal(err)
		}
		// 发送内容到服务端
		if _, err = udpConn.Write(buf[:rn]); err != nil {
			log.Fatal(err)
		}
		log.Println(i)
		i++
		time.Sleep(time.Second)
	}

	// 文件发送完成
	log.Println("file send complete.")
	time.Sleep(time.Second * 2)
}
