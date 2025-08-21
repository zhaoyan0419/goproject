package main

import "fmt"

// 执行到defer时，后边的语句不会立即执行，而是压如一个栈中，等待程序执行到最后，再从栈中取出执行
// 栈是先进后出
// 最终执行main函数，输出顺序应该是，先输出add中的sum 然后按照先入后出的顺序执行两个defer，然后输出main函数中的add返回值
// 最后按照顺序执行main中的两个defer
// 如果再add函数中，两个defer中引用了num1 和num2，并且在defer之后修改了两个变量的值，最后执行defer时候，值是不变的。相当于压入栈时候，值为多少，输出时候值就是多少
func add(num1 int, num2 int) int {
	defer fmt.Println("num1 = ", num1)
	defer fmt.Println("num2 = ", num2)
	num1 += 100
	num2 += 100
	sum := num1 + num2
	fmt.Println("sum = ", sum)
	return sum
}
func main() {
	defer fmt.Println("this is the first fmt in main")
	defer fmt.Println("this is the second fmt in main")
	fmt.Println(add(20, 10))
}
