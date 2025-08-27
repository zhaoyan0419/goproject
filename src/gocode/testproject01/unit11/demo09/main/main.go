package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个int类型管道
	c1 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 10
	}()

	// 定义一个string类型管道
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "hello i am string channel"
	}()
	//fmt.Println(<-c1) // 本身取数据就是阻塞的，等待
	// select用于选择一个已经准备好的channl来进行操作
	// 例如上边两个协程，哪个先执行好，对chan先进行插入值结束，select就选择哪个
	// case后边只能接io操作，在上边两个协程中，在赋值之前都在阻塞，谁先准备好，select就执行谁
	select {
	case v := <-c1:
		fmt.Println("intchan", v)
	case value := <-c2:
		fmt.Println("stringchan", value)
	default:
		fmt.Println("所有chan都在阻塞")
	}

}
