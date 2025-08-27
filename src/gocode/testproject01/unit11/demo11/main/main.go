package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 输出数字
func printNum() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 除法操作，使用defer+recover来捕获错误
func devide() {
	defer wg.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经被捕捉")
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0

	result := num1 / num2
	fmt.Printf("num1 / num2 = %.5f\n", result)
	// 在result时候，就已经panic了，下边代码都不会被执行
	result2 := num1 / num2
	fmt.Println(result2)
}

// 除法操作，自定义错误
func devide2() {
	defer wg.Done()
	num1 := 20
	num2 := 0
	if num2 == 0 {
		fmt.Println("错误！ 错误！ 错误！ 除数为0，请检查输入")
		return
	} else {
		result := num1 / num2
		fmt.Println(result)
	}
	fmt.Println("因为没有return，执行完if又继续向下执行了")

}
func main() {

	wg.Add(2)
	go printNum()
	//go devide2()
	go devide()
	wg.Wait()
}
