package main

import "fmt"

// 只有类型为[3]int的数组才能传入test1函数，并且数组默认传递方式是值传递
func test1(arr [3]int) [3]int {
	arr[0] = 100
	return arr
}

func test2(arr *[3]int) {
	(*arr)[0] = 200
}

func main() {

	var arr1 = [3]int{3, 6, 9}

	fmt.Println(arr1)
	fmt.Printf("arr1数组的类型为%T\n", arr1)

	//数组长度属于数组类型的一部分
	var arr2 = [6]int{3, 6, 9, 4, 56, 8}
	fmt.Println(arr2)
	fmt.Printf("arr1数组的类型为%T\n", arr2)
	// 将arr1传入test1函数，在test1函数将下标为0的值改为100，重新输出arr1
	// 可以看到arr1的第0个值仍然是3，所以数组默认传递方式是值传递
	fmt.Println("------------------使用返回值方式修改数组元素的值")
	test1(arr1)
	fmt.Println(arr1)
	//可以函数返回值重新赋值给原来的数组，以达到修改数组的目的
	test1_arr1 := test1(arr1)
	fmt.Println(test1_arr1)
	fmt.Println("---------------------test2使用指针方式修改数组元素的值")
	test2(&arr1)
	fmt.Println(arr1)

}
