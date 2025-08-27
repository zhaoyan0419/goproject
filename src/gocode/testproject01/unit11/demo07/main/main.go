package main

import (
	"fmt"
	"sync"
	"time"
)

// 写
func WriteData(c chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		c <- i
		fmt.Println("写入的数据为", i)
		time.Sleep(time.Second * 2)
	}
	close(c)
}

// 读
func ReadData(c chan int) {
	defer wg.Done()
	for value := range c {
		fmt.Println("读取的数据为", value)
		time.Sleep(time.Second * 3)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	var c1 chan int
	c1 = make(chan int, 50)
	go WriteData(c1)
	go ReadData(c1)
	wg.Wait()
}
