package main

import (
	"fmt"
	"strconv"
)

func changestring() {
	//讲int8 float32 bool type转换为string
	var (
		e1 int8    = 123
		e2 float32 = 3.14
		e3 bool    = false
		e4 byte    = 'a'
	)
	s1 := fmt.Sprintf("%d\n", e1)
	s2 := fmt.Sprintf("%f\n", e2)
	s3 := fmt.Sprintf("%t\n", e3)
	s4 := fmt.Sprintf("%c\n", e4)
	fmt.Printf("the value of s1 is %s,the type of s1 is %T\n", s1, s1)
	fmt.Printf("the value of s2 is %s,the type of s2 is %T\n", s2, s2)
	fmt.Printf("the value of s3 is %s,the type of s3 is %T\n", s3, s3)
	fmt.Printf("the value of s4 is %s,the type of s4 is %T\n", s4, s4)
}
func usestrconv() {
	//使用strconv进行字符串转换
	var (
		e1 int8 = 123
		//e2 float32 = 3.14
		//e3 bool    = false
		//e4 byte    = 'a'
	)
	s1 := strconv.FormatInt(int64(e1), 10)
	fmt.Println(s1)
	fmt.Printf("the value of s1 is %s,the type of s1 is %T\n", s1, s1)
}

func strtocommon() {
	//讲string类型转换为通用数据类型
	//string -> bool
	s1 := "false"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("the type of b is %T\n", b)

	//string -> int
	s2 := "123"
	b1, _ := strconv.ParseInt(s2, 10, 64)
	fmt.Printf("the type of b1 is %T\n", b1)

	s3 := "golang"
	var b2 bool
	b2, _ = strconv.ParseBool(s3)
	fmt.Printf("the value of b2 is %v,the type -f b2 is %T\n", b2, b2)
}

func ptrtest() {
	//指针操作，*为取值，&为取地址
	s1 := 123
	s1 += 123 + 100
	var ptr *int = &s1
	ptr1 := &ptr
	fmt.Println(s1)
	fmt.Println(ptr)
	fmt.Println(*ptr1)

	//fmt.Println(*ptr)
	//fmt.Println(ptr)
	//fmt.Println(*ptr1)
	//fmt.Printf("%T", ptr1)
}

func main() {
	ptrtest()
}
