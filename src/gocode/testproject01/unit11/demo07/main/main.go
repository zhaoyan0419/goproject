package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func raedChan(c chan int) {
	defer wg.Done()

	for i := range c {
		println("管道中传来了一个数字", i)
	}
}

func main() {
	wg.Add(1)
	c1 := make(chan int, 50)
	go raedChan(c1)

	for {
		var num int
		fmt.Println("请输入一个数字")
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("输入错误")
		}
		c1 <- num
	}

	wg.Wait()

}
