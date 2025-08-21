package main

import "fmt"

type A struct {
	a int
	b string
}
type B struct {
	c int
	d string
	a int
}
type C struct {
	A
	B
	int // 匿名字段也可以是基本数据类型
}

// 结构体匿名嵌套也可以是结构体的指针
type D struct {
	A
	*B
}

// 结构体嵌套，非匿名方式
type E struct {
	xxx A
}

func (a *A) Test() {
	fmt.Println("我是结构体A的方法")
}

func main() {
	c1 := C{A{10, "aaa"}, B{20, "bbb", 50}, 888}
	c1.A.a = 989
	fmt.Println(c1)
	fmt.Println(c1.A)
	fmt.Println(c1.B)
	// 因为A结构体和B结构体中都有a字段，所以不能用下边这个简短写法
	//fmt.Println(c1.a)
	// 如果结构体有相同字段，需要用完整的路径访问字段值，如下
	fmt.Println(c1.A.a)
	fmt.Println(c1.B.a)
	// 输出匿名字段的值
	fmt.Println(c1.int)
	fmt.Println("--------------------------")
	// 结构体匿名嵌套也可以是指针
	d1 := D{A{100, "zhaoyan"}, &B{10, "hezhenyu", 20}}
	fmt.Println(d1)
	fmt.Println(*d1.B)
	fmt.Println("------------------------")
	e1 := E{A{999, "zhoushuang"}}
	fmt.Println(e1)
	fmt.Println(e1.xxx.a)
	e1.xxx.Test()

}
