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

func DialToServer() {
	DialAddr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", DialAddr)
	if err != nil {
		log.Fatal("Dial to Server Failed", err)
	}
	defer conn.Close()
	log.Println("Dial to Server Success，Server Addr: ", conn.RemoteAddr().String())
	wg := sync.WaitGroup{}
	wg.Add(1)
	ClientReadPing(conn, &wg)
	wg.Wait()
}

func ClientReadPing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message := HBMsg{}
		decoder := gob.NewDecoder(conn)
		err := decoder.Decode(&message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("The Conn Was Closed")
				break
			}

			log.Println(err)
			continue
		}
		if message.Code == "PING" {
			log.Printf("获取到了Server发来的心跳，心跳ID：%d，即将发送PONG来相应Server", message.ID)
			ClientWritePong(conn, message)
		}
	}
}

func ClientWritePong(conn net.Conn, msg HBMsg) {
	pongMsg := HBMsg{
		ID:      rand.Uint(),
		Code:    "PONG",
		Content: fmt.Sprintf("本PONG用来回复ID为%d的PING", msg.ID),
		Time:    time.Time{},
	}
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(pongMsg)
	if err != nil {
		log.Println("PONG发送失败", err)
		return
	}
	log.Println("客户端回复了PONG", conn.RemoteAddr().String(), pongMsg.Content)
	return
}
