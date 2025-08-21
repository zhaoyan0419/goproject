package main

import "fmt"

func main() {
	var arr1 [3]int16
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("arr[0]的地址是", &arr1[0], "默认值为", arr1[0])
	fmt.Println("arr[1]的地址是", &arr1[1], "默认值为", arr1[1])
	fmt.Println("arr[2]的地址是", &arr1[2], "默认值为", arr1[2])
	fmt.Println("arr的值是", arr1)
	fmt.Printf("arr的值是%v\n", arr1)
	fmt.Printf("arr的地址是%p\n", &arr1)
}
