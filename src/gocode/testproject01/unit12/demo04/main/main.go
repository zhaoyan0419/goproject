package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	ReflectValue := reflect.ValueOf(i)
	// 因为传进来的i是个指针，对指针进行setint不合理，需要先使用elem方法，将指针转化成地址里的值在进行setint修改
	//ReflectValue.Elem().SetInt(40)
	ReflectValue.Elem().SetString("heiheihei")
}
func main() {
	// 修改基本数据类型的值
	//var num1 int = 100
	num2 := "hahaha"
	// 想在函数中修改main中变量的值，需要向函数中传入指针
	testReflect(&num2)
	fmt.Println(num2)
}
