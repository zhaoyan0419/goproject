package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "test.txt"
	destFile := "test333.txt"
	srcContent, srcErr := os.Open(srcFile)
	if srcErr != nil {
		fmt.Println("源文件打开失败：", srcErr)
		return
	}
	defer srcContent.Close()
	destContent, destErr := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if destErr != nil {
		fmt.Println("目标文件打开失败", destErr)
		return
	}
	defer destContent.Close()
	srcReader := bufio.NewReader(srcContent)
	destWriter := bufio.NewWriter(destContent)
	//destWriter := destContent.WriteString()

	for {
		str, err := srcReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取结束", err)
				break
			}

		}
		destWriter.WriteString(str)
	}
	destWriter.Flush()

}
