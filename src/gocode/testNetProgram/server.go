package testNetProgram

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

type HBMsg struct {
	ID      uint      `json:"ID,omitempty"`
	Code    string    `json:"code,omitempty"`
	Content string    `json:"content,omitempty"`
	Time    time.Time `json:"time,omitempty"`
}

func ListenTcp() {
	ServerAddr := ":9999"
	listener, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		log.Fatal("Server Listen Failed", err)
	}
	log.Println("Server Listen Success,ServerAddress:", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Listener Accept Failed", err)
			continue
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	log.Println("Conn Set Success,The Remote Addr:", conn.RemoteAddr().String())
	// 向客户端发送心跳
	wg := sync.WaitGroup{}
	wg.Add(1)
	go HeartBeatToClient(conn, &wg)
	wg.Wait()

}

func HeartBeatToClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	const maxPingNum = 3
	pingErrCounter := 0
	ticker := time.NewTicker(time.Second * 2)
	for t := range ticker.C {
		pingMsg := HBMsg{
			ID:      rand.Uint(),
			Code:    "PING",
			Content: "",
			Time:    t,
		}
		encoder := gob.NewEncoder(conn)
		err := encoder.Encode(pingMsg)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("连接已经断开")
				return
			}
			log.Println("向客户端写入数据失败", err)
			pingErrCounter++
			if pingErrCounter == maxPingNum {
				fmt.Println("三次心跳检测失败，conn即将被关闭")
				return
			}
			continue

		}
		log.Printf("向客户端发送心跳，心跳ID: %d,发送远程地址: %s\n", pingMsg.ID, conn.RemoteAddr().String())
	}

}
