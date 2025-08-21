package main

import "fmt"

func testSlice() {
	// 使用普通for循环对切片进行遍历
	slice1 := make([]int, 4, 20)
	slice1[0] = 1
	slice1[1] = 4
	slice1[1] = 7
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice2[%v] = %v\n", i, slice1[i])
	}
	fmt.Println("-------------------------------")
	// 使用for range对切片进行遍历
	for k, v := range slice1 {
		fmt.Printf("slice2[%v] = %v\n", k, v)
	}
}

// 切片使用要么使用make初始化，要么给定元素具体数值，否则不可以直接使用
func testSlice2() {
	var slice1 []int
	fmt.Println(slice1)
}

// 切片下标不能越界
func testSlice3() {
	var testArr1 = [6]int{1, 4, 7, 2, 5, 8}
	slice1 := testArr1[1:4]
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	fmt.Println(slice1[2])
	fmt.Println(slice1[3])
}

// 对切片再进行切片
func testSlice4() {
	slice1 := make([]int, 4, 20)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 5
	}
	fmt.Println("slice1 = ", slice1)
	slice2 := slice1[1:]
	fmt.Println("slice2 = ", slice2)
	fmt.Println("对slice1中元素进行修改")
	slice1[1] = 1000
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
func testSlice5() {
	var testArr1 = [6]int{1, 2, 3, 4, 5, 6}
	slice1 := testArr1[1:4]
	fmt.Println(slice1)
	fmt.Println("slice1的长度", len(slice1))
	fmt.Println("slice1的容量", cap(slice1))
	fmt.Printf("slice1的内存地址是%p\n", &slice1[0])
	fmt.Printf("testArr1[1]的内存地址是%p\n", &testArr1[1])
	// 对slice中追加元素
	// 底层原理：
	// 底层追加元素的时候对数组进行扩容，老数组扩容为新数组
	// 创建一个新数组，将老数组中的2 3 4 赋值到新数组中，在新数组中追加1å
	slice2 := append(slice1, 1, 1, 1, 1, 1)
	fmt.Println(slice2)
	slice2[0] = 100
	fmt.Println(slice1, slice2)
	fmt.Println(testArr1)
}

// 切片后边可以追加一个切片
func testSlice6() {
	slice1 := []int{2, 4, 6, 8, 10}
	slice2 := []int{100, 200, 300}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1, slice2)
}

// 切片的拷贝
func testSlice7() {
	var slice = []int{1, 4, 7, 3, 6, 9}
	var slice2 = make([]int, 10)
	copy(slice2, slice)
	fmt.Println(slice)
	fmt.Println(slice2)
	slice2[0] = 100
	fmt.Println(slice)
	fmt.Println(slice2)
}

func testSlice8() {
	arr1 := [6]int{100, 200, 300, 400, 500, 600}
	slice1 := arr1[1:4]
	fmt.Println("arr1的元素个数为", len(arr1))
	fmt.Println("输出数组arr1 = ", arr1)
	fmt.Println("输出slice1 = ", slice1)
	slice2 := append(slice1, 5, 6)
	fmt.Println(arr1, slice2)
	slice3 := append(slice1, 555, 666, 777, 888)
	fmt.Println(slice3, arr1)
}
func main() {
	testSlice8()
}
