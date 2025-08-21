package main

import (
	"fmt"
	"os"
)

// err := os.Rename("newTest.txt", "test.txt")
// //os.Rename("test.txt", "newTest.txt")
//
//	if err != nil {
//		fmt.Println("文件改名出错", err)
//	} else {
//
//		fmt.Println("文件修改名称完成")
//	}
func main() {
	// 下边的操作不需要open 和close，因为已经被封装到内部了
	file := "./test.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("文件读取出错，错误为：", err)
		return
	}
	// 如果读取成功
	fmt.Println(string(content))
}
