package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	// 方式一：必须按照顺序
	var s Student = Student{"zs", 30}
	fmt.Println(s)
	// 方式二
	s1 := Student{Name: "yanyan", Age: 28}
	fmt.Println(s1)
	// 方式三:返回指针类型，使用变量接收
	var s2 *Student = &Student{"hzy", 27}
	fmt.Println(*s2)
}
