package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile1() {
	srcFile := "./test.txt"
	destFile := "./test111.txt"
	readContent, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	// 当函数退出时，让file关闭，防止内存泄漏
	defer readContent.Close()
	//
	writeContent, err1 := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("文件打开失败", err1)
		return
	}
	defer writeContent.Close()

	// 写文件操作
	// 创建一个写缓冲区
	writer := bufio.NewWriter(writeContent)

	// 创建流
	reader1 := bufio.NewReader(readContent)
	//读取操作
	for {
		str, err1 := reader1.ReadString('\n')

		if err1 == io.EOF {

			break
		}
		// 调用writer的WriteString方法
		writer.WriteString(str)
		// 流在缓冲区，需要刷新数据到文件中

	}
	writer.Flush()
}

func main() {
	srcFile := "test.txt"
	destFile := "test222.txt"
	//srcContent, srcErr := os.Open(srcFile)
	//if srcErr != nil {
	//	fmt.Println("src文件打开失败", srcErr)
	//	return
	//}
 
	// 读取文件
	srcContent, srcErr := os.ReadFile(srcFile)
	if srcErr != nil {
		fmt.Println("文件读区读取失败：", srcErr)
		return
	}

	// 写出文件
	destErr := os.WriteFile(destFile, srcContent, 0666)
	if destErr != nil {
		fmt.Println("写出失败", destErr)
	}
}
