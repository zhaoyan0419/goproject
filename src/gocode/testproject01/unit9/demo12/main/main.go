package main

import (
	"fmt"
	"os"
	"sync"
)

type AInterface interface {
	a()
	BInterface
	CInterface
}
type BInterface interface {
	b()
}

type CInterface interface {
	c()
}

type Stu struct {
	Name string
	Age  int
}

type Person struct {
	sync.Mutex
}

func (s Stu) a() {
	fmt.Println("我是Stu的a()方法")
}

func (s Stu) b() {
	fmt.Println("我是Stu的b()方法")
}

func (s Stu) c() {
	fmt.Println("我是Stu的c()方法")
}
func (p Person) a() {
	fmt.Println("我是Person的a方法")
}

func goAInterface(a AInterface) {
	a.b()

}

func main() {
	s1 := Stu{"zhaoyan", 28}
	goAInterface(s1)
	//p1 := Person{}
	// 因为person只实现了AInterface中的一个a方法（相当于person并没有实现AInterface接口），没有实现其他方法么所以p1无法传入goAinterface
	//goAInterface(p1)
	fmt.Println(os.Getenv("PATH"))
}
