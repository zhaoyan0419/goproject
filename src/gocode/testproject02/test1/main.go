package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const udp = "udp"

func UDPReceiverMulticast() {
	// 1组播监听地址
	ifi, err := net.InterfaceByName("ens160")
	if err != nil {
		log.Fatal(err)
	}
	address := "224.1.1.2:6789"
	gaddr, err := net.ResolveUDPAddr(udp, address)
	if err != nil {
		log.Fatal(err)
	}
	// 2 组播监听
	UdpConn, err := net.ListenMulticastUDP(udp, ifi, gaddr)
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

// 广播发送端
func UDPSenderBroadcast() {
	raddr, err := net.ResolveUDPAddr(udp, "192.168.1.255:9999")
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
		data := []byte("向广播地址发送数据。。。" + string(i))
		_, err := udpConn.WriteToUDP(data, raddr)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		log.Println("数据发送成功，第", i, "次发送")
		i++
		time.Sleep(time.Second)
	}

}

func main() {
	UDPSenderBroadcast()
}
