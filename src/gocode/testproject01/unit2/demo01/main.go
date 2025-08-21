package main

import (
	"fmt"
	"unsafe"
)

func main() {
	age := -1233333
	fmt.Printf("the type of age is %T \n", age)
	fmt.Println(unsafe.Sizeof(age))

	var num1 float32 = 3.14
	fmt.Println(num1)
	fmt.Printf("num1 type is %T", num1)

	var num2 float32 = 314e-2
	fmt.Println(num2)

	num3 := 3.14
	fmt.Printf("num3 type is %T", num3)
	fmt.Println()

	var num4 int = '中'
	fmt.Printf("the type of num4 is %T \n,the charactor is %c \n,the num is %d", num4, num4, num4)
	fmt.Println()

	zhen := true
	fmt.Printf("the type of zhen is %T\n", zhen)

	string1 := "aaa"
	fmt.Println(string1)
	string1 = "bbb"
	fmt.Println(string1)
	string2 := `
kind: Pod
apiVersion: v1
metadata: 
  name: test1
  namespace: default
  labels:
    a: b
spec:
  containers:
  - name: c1
    image: nginx:latest
    imagePullPolicy: IfNotPresent
`
	fmt.Println(string2)
	s1 := "abc" + "def"
	fmt.Println(s1)
	s1 += "hijk"
	fmt.Println(s1)

	var s2 = "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def"
	fmt.Println(s2)

	//各种数据类型的默认值
	var (
		a int
		b float64
		c bool
		d string
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//类型转换
	var n1 int = 123
	var n2 float32
	n2 = float32(n1)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Printf("the type of n1 is %T\n", n1)
	fmt.Printf("the type of n2 is %T\n", n2)

	//int8转换为int64会有精度的丢失
	var q1 int64 = 88888
	var q2 int8 = int8(q1)
	fmt.Println(q1)
	fmt.Println(q2)

	//进行算数运算时候，需要先转类型
	var w1 int32 = 123
	var w2 int64 = int64(w1) + 10
	fmt.Println(w1)
	fmt.Println(w2)

}
