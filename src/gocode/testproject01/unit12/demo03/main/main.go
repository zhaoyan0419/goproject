package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func testReflect(i interface{}) {
	tp := reflect.TypeOf(i)
	vl := reflect.ValueOf(i)

	// 获取变量的类别:
	// 方式以，使用reflect.type.kind
	k1 := tp.Kind()
	fmt.Println("方式一使用reflect.value.kind返回kind为:", k1)
	// 方式二，以reflect.value.kind
	k2 := vl.Kind()
	fmt.Println("方式二使用reflect.value.kind返回kind为:", k2)
	// 获取变量的类型
	value := vl.Interface()
	v, ok := value.(Person)
	if ok {
		fmt.Printf("结构体的类型是%T\n", v)
	}
}
func main() {
	p1 := Person{
		Name: "zhaoyan",
		Age:  28,
	}

	testReflect(p1)
}
