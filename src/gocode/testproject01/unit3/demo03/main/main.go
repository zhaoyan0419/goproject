package main

import (
	"fmt"
	"github.com/src/gocode/testproject01/unit3/demo03/tuil"
)

func init() {
	fmt.Println("init函数被执行")
}

var num1 int = test()

func test() int {
	fmt.Println("当给num1赋值时候，赋的是test函数的返回值，我现在这条输出就是test函数的一部分")
	return 10
}

func main() {
	fmt.Println("main函数被执行")
	fmt.Printf("name= %s,age = %d,sex = %s", tuil.Name, tuil.Age, tuil.Sex)
}
