package main

import "fmt"

type Persion struct {
	name string
}

func (p Persion) Test() {
	p.name = "lisi"
	fmt.Println("方法中没有使用指针", p.name)
}

func (p *Persion) Test2() {
	p.name = "wangwu"
	fmt.Println("方法中使用指针修改结构体里的值后name为", p.name)
}

type myInt int

func (m *myInt) myIntTest() {
	*m = 1000
	fmt.Println(*m)
}

func main() {
	p1 := Persion{}
	p1.name = "zhangsan"
	fmt.Println("最初始结构体中的name为", p1.name)
	fmt.Printf("p1的地址是%p\n", &p1)
	p1.Test()
	p1.Test2()
	(&p1).Test2()
	fmt.Println(p1.name)
	var m1 myInt = 100
	m1.myIntTest()
}
