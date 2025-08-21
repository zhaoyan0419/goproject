package main

import "fmt"

func test() {
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
	//  panic: runtime error: integer divide by zero
	//
	//	goroutine 1 [running]:
	//	main.test()
	//	/Users/zhaoyan/goproject/src/gocode/testproject01/unit4/demo01/main/main.go:8 +0x1c
	//	main.main()
	//	/Users/zhaoyan/goproject/src/gocode/testproject01/unit4/demo01/main/main.go:13 +0x1c

}

// defer + recover机制处理错误
func test1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经被捕获")
			fmt.Println("err为", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)

}
func main() {

	test1()
	test1()
	test1()
	test1()
	fmt.Println("上边的函数执行成功")
}
