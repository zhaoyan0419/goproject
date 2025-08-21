package main

import "fmt"

// 函数的定义
func add1() {
	num1 := 20
	num2 := 10
	sum := 0
	sum += num1
	sum += num2
	fmt.Println(sum)
}
func add2(a int, b int) (int, int) {

	cal, sum := 0, 0
	sum += a
	sum += b
	cal = a - b
	return sum, cal
}

func exchangeNum(num1 int, num2 int) (int, int) {
	t := 0
	t = num1
	num1 = num2
	num2 = t
	return num1, num2
}

// 可变参数数量,可变参数会以切片传进函数
func add3(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 基本数据类型和数组都是值传递，函数中对参数进行操作不会修改主函数中参数的值
// 使用指针，可以将参数地址传入函数，将参数对应地址的值进行修改从而修改主函数中参数的值
func test1(num *int) {
	*num = 30

	//num := 10
	//fmt.Println(num)
	//test1(&num)
	//fmt.Println(num)
}

func test2(a int) {
	fmt.Println(a)
}
func test3(num1 int, num2 int, c func(int)) {
	c(num2)
	//a := test3
	//test3(100, 200, test2)
	//a(500, 600, test2)
}

// 对返回值进行命名
func test4(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	fmt.Println(test4(300, 200))
}
