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
	// 调用typeof
	tp := reflect.TypeOf(i)
	fmt.Println(tp)
	fmt.Printf("您传入的参数反射后类型是%T\n", tp)
	// 调用valueof
	vl := reflect.ValueOf(i)
	fmt.Println(vl)
	// 虽然输出显示的是一个person的实例，实际上并不是。无法使用person.name获取值
	fmt.Printf("您传入的参数反射后值是%T\n", vl)

	// 反方向转
	i2 := vl.Interface()
	value, ok := i2.(Person)
	if ok {
		fmt.Println("断言成功后，姓名是", value.Name)
		fmt.Println("断言成功后，年龄是", value.Age)
	}

}

func main() {
	p1 := Person{
		Name: "zhaoyan",
		Age:  28,
	}
	testReflect(p1)
}
