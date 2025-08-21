package main

import (
	"fmt"
)

// for 循环
func test1() {
	sum := 0
	for i := 1; i <= 5; i += 1 {
		sum += i
	}
	fmt.Println(sum)

}

// for死循环
func test2() {
	for {
		fmt.Println("你好，死循环了")
	}
}

// for range 循环
func test3() {
	str := "hello world你好"
	for i, j := range str {
		fmt.Printf("str[%v]:%c\n", i, j)
	}
}
func test4() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			break
		}

	}
}

func test5() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d   ", j, i, i*j)
		}
		fmt.Println()
	}
}

// break，直接退出当前循环
func test6() {
	//可以给循环加标签la，满足条件后，break标签循环
la:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == 2 && j == 2 {
				break la
			}
		}
	}
}

func test7() {
	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}

	}
}

// continue结束本次循环，继续下一次循环，continue和break一样，可以使用label
func test8() {
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
func main() {
	test6()
}
