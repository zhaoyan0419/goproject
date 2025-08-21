package main

import (
	"fmt"
	"os"
)

func main() {
	f1, err := os.Open("./test.txt")
	// 如果出错，输出一下错误
	if err != nil {
		fmt.Println("文件打开出错", err)

	}
	fmt.Printf("文件 = %v", f1)
	CloseErr := f1.Close()
	if CloseErr != nil {
		fmt.Println("文件关闭失败")
	}

}
