package main

import "fmt"

// 算数运算符
func calculate1() {
	// +
	var a, b int = 20, 10
	fmt.Println(a + b)
	fmt.Println("avc" + "def")
	// /除以
	c := a / b
	var d float64
	d = 10.0 / 3
	fmt.Println(c, d)
	fmt.Printf("%T", d)
	fmt.Println()
	// 取模% 取模的符号只和被除数符号相关
	fmt.Println(10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % 3)
	fmt.Println(-10 % -3)
	// ++ 自增 -- 自减 每次+1 或者 -1
	e := 10
	e++
	fmt.Println(e)
	f := 10
	f--
	fmt.Println(f)
}

// 赋值运算符
func calculate2() {
	// =  += -= *= /= %=
	var a int = 10
	fmt.Println(a)
	var b int = 20
	a += b
	fmt.Printf("a = %d\n", a)

	//交换值
	c := 20
	d := 10
	fmt.Printf("c = %v,d = %v\n", c, d)
	t := c
	c = d
	d = t
	fmt.Printf("c = %v,d = %v\n", c, d)
}

// 关系运算符
func calculate3() {
	// == != > < >= <=
	fmt.Println(3 <= 5)
}

// 逻辑运算符
func calculate4() {
	// &逻辑与     ||逻辑或      !逻辑非

	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	fmt.Println(!b)
}
func main() {
	calculate4()
}
