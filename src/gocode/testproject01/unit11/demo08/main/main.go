package main

import "fmt"

func main() {
	// 默认情况下，管道是双向的，可读可写
	//var c1 chan int
	//c1 = make(chan int, 10)
	// 声明为只写
	var c2 chan<- int // 管道具备 <-   只写性质
	c2 = make(chan<- int, 3)
	c2 <- 10
	//num1 := <-c2 // 此处会报错，因为c2只写，无法读取
	//fmt.Println("c2", c2)

	//声明为只读
	var c3 <-chan int
	fmt.Println(c3 == nil)
	//只有当 c3 不为空时候，才能在c3中读取数据赋值给num1
	if c3 != nil {
		num1 := <-c3
		fmt.Println(num1)
	}
	// 在初始化时候给只读channel赋值
	c4 := make(chan int, 5)
	c4 <- 10
	c4 <- 20
	c4 <- 30
	c4 <- 40
	c4 <- 50
	var c5 <-chan int = c4
	num1 := <-c5
	fmt.Println("num1 = ", num1)
}
