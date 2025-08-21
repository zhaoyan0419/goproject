package main

import "fmt"

// 接口的定义
type SayHello interface {
	//声明方法，暂时还没有实现的方法
	SayHello()
}

// 接口的实现
type Chinese struct {
	Name string
	Age  int
}
type American struct {
	Name string
	Age  int
}
type Japanese struct {
	Name string
	Age  int
}

// Chinese结构体实现Sayhello这个接口

func (c Chinese) SayHello() {
	fmt.Println("你好")
}

// American结构体实现SayHello这个接口
func (a American) SayHello() {
	fmt.Println("Hello")
}

// 定义一个函数：专门用来个过人打招呼的函数，接收具备SayHello接口的能力的变量
func greet(s SayHello) {
	s.SayHello()
}

// 只要是自定义的数据类型都可以实现接口
type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("自定义的数据类型也可以实现接口")
}

func main() {
	c1 := Chinese{"zhaoyan", 28}
	a1 := American{"bryant", 45}
	j1 := Japanese{"hzy", 27}
	fmt.Println(j1)
	c1.SayHello()
	a1.SayHello()
	greet(c1)
	greet(a1)
	// 接口本身不能创建实例,但是可以来接收已经实现了这个接口的结构体实例
	var s1 SayHello = Chinese{"hehe", 20}
	s1.SayHello()
	fmt.Println("我的类型是myint，验证自定义类型实现方法")
	var m1 MyInt = 100
	greet(m1)

}
