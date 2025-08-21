package main

import "fmt"

type Student struct {
	age int
}

type Persion struct {
	age int
}

func main() {
	// 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字，个数，类型）
	// 结构体的转换必须使用强制类型转换
	s1 := Student{age: 100}
	p1 := Persion{}
	fmt.Println(s1, p1)
	p1 = Persion(s1)
	fmt.Println(p1)
}
