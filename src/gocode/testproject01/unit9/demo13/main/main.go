package main

import (
	"fmt"
)

type AInterface interface {
	a()
}

type Stu struct {
	Name string
	Age  int
}
type Teacher struct {
	Name string
	Age  int
}

func (s Stu) a() {
	fmt.Println("我是学生")
}
func (t Teacher) a() {
	fmt.Println("我是老师")
}

func goAInterface(a AInterface) {
	a.a()

}

func main() {
	s1 := Stu{"zhaoyan", 20}
	t1 := Teacher{"zhoushuang", 30}
	// 定义一个AInterface接口的切片，接收具备AInterface接口能力的变量
	slice1 := []AInterface{s1, t1}
	for _, v := range slice1 {
		v.a()
	}
}
