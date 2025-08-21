package main

import "fmt"

func main() {
	// 数组内存分析，数组的地址，和数组中第一个元素的地址是相同的
	// 二维数组内存分析，二维数组的地址和arr1[0]的地址是相同的，并且和arr1[0][0]也是相同的
	arr1 := [2][3]int{}
	fmt.Println(arr1)
	arr1[0][1] = 9
	fmt.Println(arr1)
	fmt.Printf("arr1的地址为%p\n", &arr1)
	fmt.Printf("arr1[0]的地址为%p\n", &arr1[0])
	fmt.Printf("arr1[0][0]的地址为%p\n", &arr1[0][0])
	fmt.Printf("arr1[0][0]的地址为%p\n", &arr1[1][0])
	//[[0 0 0] [0 0 0]]
	//	[[0 9 0] [0 0 0]]
	//	arr1的地址为0x140000b8000
	//	arr1[0]的地址为0x140000b8000
	//	arr1[0][0]的地址为0x140000b8000
	//	arr1[1][0]的地址为0x140000b8018

	fmt.Println("---------------------二维数组的遍历")
	arr2 := [3][3]int{{1, 6, 8}, {2, 5, 7}, {9, 5, 3}}

	fmt.Println(arr2[1])
	for k, v := range arr2 {
		for i, j := range v {
			fmt.Printf("arr1[%d][%d]的值为%d\t", k, i, j)
		}
		fmt.Println()
	}
	fmt.Println("--------------------------二维数组遍历方式二")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]:\t", i)
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "\t")
		}
		fmt.Println()
	}
}
