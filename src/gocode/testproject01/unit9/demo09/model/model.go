package model

import "fmt"

type Animal struct {
	Age    int
	Weight int
}

// Animal的两个方法。

func (a *Animal) Shout() {
	fmt.Println("我可以大声喊叫")
}
func (a *Animal) ShowInfo() {
	fmt.Printf("动物的年龄是：%v，动物的体重是：%v\n", a.Age, a.Weight)
}

// 为了复用性，加入匿名结构体
// 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母答谢或者小些的字段/方法都可以使用
type Cat struct {
	Animal
	Age int // 如果结构体中和嵌套的结构体中有相同的字段，获取字段值时候会采用就近原则
}
type Dog struct {
	Animal
}

func (c *Cat) Scrath() {
	fmt.Println("我是小猫，我可以挠人")
}
