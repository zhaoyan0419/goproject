package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写操作
func writeData(c chan int) {
	defer wg.Done()
	// chan的长度为5，但是写入6个数，如果只写不读就会出现死锁
	for i := 1; i <= 6; i++ {
		c <- i
		fmt.Println("写入数据为", i)
		//time.Sleep(time.Second * 1)
	}
	//close(c)
}

// 读操作
func readData(c chan int) {
	defer wg.Done()
	for value := range c {
		fmt.Println("读取数据为", value)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var c1 chan int = make(chan int, 5)
	wg.Add(1)
	// 当只写不读时候，并且chan的缓冲取被写满，就会出现阻塞，死锁
	go writeData(c1)
	//go readData(c1)
	wg.Wait()

}
