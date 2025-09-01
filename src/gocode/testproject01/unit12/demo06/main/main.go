package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func changeStructField(i interface{}) {
	name1 := "wangxiaoge"
	value := reflect.ValueOf(i)
	fmt.Println(value.Elem().Field(1))
	// 修改name字段的值
	value.Elem().Field(0).SetString(name1)
	// 修改age字段的值
	value.Elem().Field(1).SetInt(100)

}
func main() {
	p1 := Student{
		Name: "zhaoyan",
		Age:  28,
	}
	// 传入p1的指针，在函数内修改p1
	changeStructField(&p1)
	fmt.Println(p1)
}
