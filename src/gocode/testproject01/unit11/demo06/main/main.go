package main

import "fmt"

func main() {
	var c1 chan int
	c1 = make(chan int, 100)
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	// 遍历前如果没有关闭管道，就会出现deadlock，所以在遍历管道之前需要将管道关闭3
	close(c1)
	for v := range c1 {
		fmt.Println(v)
	}

}
