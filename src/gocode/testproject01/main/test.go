package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func test1() {
	age := 25
	fmt.Println("age = ", age)
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	for a, b := range s1 {
		fmt.Printf("index = %d, value = %d\n", a, b)
	}
}

func test2() { //整数类型测试
	var age int16 = 286
	fmt.Printf("your age is %d\n", age)
	fmt.Printf("age 的类型是：%T\n", age)
	fmt.Println(unsafe.Sizeof(age))
}

func test3() {
	var weight float32 = 75.123
	fmt.Printf("你的体重是：%f\n", weight)
}

func test4() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := User{Name: "Alice", Age: 25}
	data, _ := json.Marshal(user) // 序列化
	fmt.Println(string(data))     // 输出: {"name":"Alice","age":25}

	jsonStr := `{"name":"Bob","age":30}`
	var u User
	_ = json.Unmarshal([]byte(jsonStr), &u) // 反序列化
	fmt.Println(u)
}

func test5() {
	var a byte = 'a'
	var b byte = '%'
	c := '('
	var d rune = '中'
	fmt.Println(a, b, c)
	fmt.Println(d)
}
func main() {
	test5()
}
