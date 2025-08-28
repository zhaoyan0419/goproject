package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	addr := "tengxunyun.zhaoyan.site:7899"
	//addr := "127.0.0.1:7899"
	con, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接失败,err:", err)
		return
	}
	defer con.Close()
	fmt.Println("连接成功", con)
	//time.Sleep(time.Second * 3)
	// 通过客户端发送但行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.stdin代表终端标准输入
	for i := 1; i <= 10; i++ {

		// 从终端读取一行用户输入的信息
		str, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("读取终端输入失败，err：", err1)
		}
		n, err2 := con.Write([]byte(str))
		if err2 != nil {
			fmt.Println("数据发送失败，err：", err2)
		}
		fmt.Printf("终端数据通过客户端发送成功，本次发送了%d字节的数据\n", n)
	}

}
