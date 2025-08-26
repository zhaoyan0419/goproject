package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello msb，我是一个协程 +" + strconv.Itoa(i))
		// 阻塞1秒
		time.Sleep(time.Second * 1)
	}
}

func main() { // 主线程
	go test()   // 开启一个协程
	go func() { // 使用匿名函数启动协程
		for i := 1; i <= 10; i++ {
			fmt.Println("我是一个匿名函数，并且是一个协程，这是我第" + strconv.Itoa(i) + "次执行")
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println("hello golang 我在main主线程中执行+" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
