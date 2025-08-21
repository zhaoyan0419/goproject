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
func (s Stu) PlayComputerGame() {
	fmt.Println("我是玩电脑游戏，我是Stu特有的方法")
}
func (t Teacher) HomeWork() {
	fmt.Println("老师要批改作业")
}
func goAInterface(a AInterface) {
	a.a()
	// 看看传进来的a是否能转成Stu类型并且赋值给s100变量
	//var s100 Stu = a.(Stu)
	//s100.PlayComputerGame()
	//if s100, ok := a.(Stu); ok {
	//	s100.PlayComputerGame() // 只有是 Stu 时才调用
	//} else {
	//	fmt.Println("传入的不是 Student 类型，无法执行Stu类型的方法PlayComputerGame")
	//}
	switch v := a.(type) {
	case Stu:
		v.PlayComputerGame()
	case Teacher:
		v.HomeWork()
	}
}

func main() {
	s1 := Stu{"zhaoyan", 20}
	t1 := Teacher{"zhoushuang", 30}
	goAInterface(s1)
	fmt.Println("--------------------下边将Teacher类型传入goAInterface函数")
	goAInterface(t1)

}
