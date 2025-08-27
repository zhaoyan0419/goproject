package main

import (
	"fmt"
	"net"
)

func main() {
	add := "tengxunyun.zhaoyan.site:9797"
	con, err := net.Dial("tcp", add)
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("连接成功", con)
}
