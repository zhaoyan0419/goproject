package main

import "fmt"

func main() {
	// 创建一个int类型管道
	var c1 chan int
	// 初始化管道可以存放3个int类型的数
	c1 = make(chan int, 3)
	fmt.Printf("验证channel是否是引用类型,%v\n", c1)
	c1 <- 10
	c1 <- 20
	c1 <- 30
	fmt.Println("长度", len(c1))
	fmt.Println("容量", cap(c1))
	fmt.Println(c1)
	// 将c1中的10取出并赋值给num1
	num1 := <-c1
	fmt.Println(num1)
	c1 <- 40
	fmt.Println(c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	// 管道中没有数据了，不能重复再拿，
	//fmt.Println(<-c1)
	// 再次向管道中插入数据
	c1 <- 50
	c1 <- 60
	c1 <- 70
	// 关闭管道，验证关闭管道后还能不能拿到管道中的数据
	close(c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(c1)
	//管道关闭后，不能再向管道中加入数据
	c1 <- 80
}
