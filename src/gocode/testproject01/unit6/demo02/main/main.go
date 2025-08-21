package main

import "fmt"

func main() {
	// 定义切片，make函数三个参数，切片类型/切片长度/切片容量
	// make底层创建一个数组，对外不可见，不可以直接操作这个数组，要通过切片去间接的访问各个元素
	slice := make([]int, 4, 20)
	fmt.Println(slice)
	fmt.Println("slice切片的长度为", len(slice))
	fmt.Println("slice切片的容量为", cap(slice))
	slice[0] = 88
	slice[1] = 66
	fmt.Println(slice)
	// 另一个方式定义切片
	slice2 := []int{1, 4, 7}
	fmt.Println(slice2)
	fmt.Println("slice切片的长度为", len(slice2))
	fmt.Println("slice切片的容量为", cap(slice2))
}
