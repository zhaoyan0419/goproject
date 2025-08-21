package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	srcFile := "./test1.txt"
	// |os.O_APPEND
	content, err1 := os.OpenFile(srcFile, os.O_RDWR|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("文件打开失败", err1)
		return
	}
	defer content.Close()

	// 写文件操作
	writer := bufio.NewWriter(content)
	writer.WriteString("你好马士兵\n")
	// 流在缓冲区，需要刷新数据到文件中
	writer.Flush()
}
