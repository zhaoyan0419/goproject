package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {
	fmt.Println("调用了Print()方法")
	fmt.Println("人的姓名是：", p.Name)
}
func (p Person) GetSum(n1, n2 int) int {
	sum := n1 + n2
	return sum
}
func (p Person) Set(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println(p.Name, p.Age)
}

func testReflect(i interface{}) {
	value := reflect.ValueOf(i)
	fmt.Println(value)
	// 通过reflect.value类型操作结构体字段
	n := value.NumField() // 获取到结构体有多少个字段，后边可以通过遍历获取每个字段的值
	fmt.Println("传入的struct的字段数:", n)
	for i := 0; i < n; i++ {
		fmt.Printf("第%d个字段的值是:%v\n", i, value.Field(i)) // value.field(index)来获取每个字段的值
	}
	// 通过reflect.value类型操作结构体内部的方法
	methodNum := value.NumMethod()
	fmt.Println("方法的数量为：", methodNum)
	// 调用方法（方法首字母必须大写，才能call）
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf("wangxiaoge"))
	args = append(args, reflect.ValueOf(35))
	value.Method(2).Call(args)
	// 调用另一个方法
	args2 := make([]reflect.Value, 0)
	args2 = append(args2, reflect.ValueOf(20))
	args2 = append(args2, reflect.ValueOf(30))
	fmt.Printf("%T\n%v", value.Method(0).Call(args2)[0].Int(), value.Method(0).Call(args2)[0].Int())
}

func main() {
	p1 := Person{
		Name: "zhaoyan",
		Age:  28,
	}
	testReflect(p1)
}
