package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型
	var num1 int = 100
	// 调用typeof函数，返回reflect.type类型数据
	tp := reflect.TypeOf(num1)
	fmt.Printf("tp的具体类型是%T\n", tp)
	fmt.Println(tp)
	fmt.Println(tp.Name())
	fmt.Println(tp.Kind())
	// 调用valueof函数，返回reflect.value类型数据
	value := reflect.ValueOf(num1)
	fmt.Printf("value的具体类型是%T\n", value)
	fmt.Println(value)
	//var num2 int = 200
	// 因为valueof返回的是value类型，不能和int类型进行运算
	//sum := value + num2
	sum := value.Int() + 80
	fmt.Println(sum)
	fmt.Println("---------------------------------验证value.interface")
	str := "haha"
	vvv := reflect.ValueOf(str)
	// 使用interface方法，返回一个interface，但是interface会保存动态类型和值
	i := vvv.Interface()
	fmt.Printf("反射得到的vvv的interface方法后，i的类型和值分别是 %T %v\n", i, i) // 打印string haha
	fmt.Println(i)
	// 断言引入
	n, ok := i.(int)
	fmt.Printf("i的类型是%T\n", i)
	fmt.Println(n, ok)
}
