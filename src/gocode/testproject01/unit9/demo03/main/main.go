package main

import "fmt"

// 方法的声明和使用
type A struct {
	name string
}

// 结构体类型是值传递
func (a *A) test() {
	a.name = "zongshan"
	fmt.Println(a.name)
}
func main() {
	var a1 A
	a1.name = "zhaoyan"
	fmt.Println(a1.name)
	a1.test()

}
