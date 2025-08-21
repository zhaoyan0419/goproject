package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "./test.txt"
	content, err := os.Open(file)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	// 当函数退出时，让file关闭，防止内存泄漏
	defer content.Close()
	// 创建流
	reader1 := bufio.NewReader(content)
	//读取操作
	for {
		str, err1 := reader1.ReadString('l')

		if err1 == io.EOF {
			fmt.Println(io.EOF)
			break
		}
		fmt.Println(str)
	}
	//str, _ := reader1.ReadString('l')
	//
	//fmt.Println(str)
}
