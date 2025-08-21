package model

import "fmt"

type person struct {
	Name string
	age  int // 其他包不能直接访问
}

// 定义工厂模式的函数，相当于构造器
func NewPerson(n string) *person {
	return &person{Name: n}
}

//定义set和get函数，对age字段进行封装，因为在函数中加一系列的操作，确保被封装字段的安全合理性

func (p *person) Setage(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("对不起，您传入的年龄不符合要求")
	}
}
func (p *person) GetAge() int {
	return p.age
}
