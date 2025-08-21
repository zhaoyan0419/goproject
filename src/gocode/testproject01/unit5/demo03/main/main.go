package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("请输入第%d个人的成绩为", i+1)
		fmt.Scanln(&arr1[i])
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("第%d个学生的成绩为%d\n", i+1, arr1[i])
	}
	fmt.Println("-------------------------------------------------------------------------------------------")
	for k, v := range arr1 {
		fmt.Printf("第%d个学生的成绩为%d\n", k+1, v)

	}
}
