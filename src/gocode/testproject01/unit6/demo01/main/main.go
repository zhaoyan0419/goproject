package main

import "fmt"

func main() {
	var arr1 = [6]int{2, 5, 7, 6, 4, 2}
	// 切片可以从数组中切一段出来，切片切出来的是左开右闭区间
	var slice1 []int = arr1[1:3]
	fmt.Println(slice1)
	fmt.Println("slice1切片的元素个数为：", len(slice1))
	fmt.Println("slice1切片的容量为：", cap(slice1))
	// 切片是从数组中切出来的片段，修改切片元素后，底层数组的元素也会变化
	slice1[0] = 89
	fmt.Println("修改后的slice是", slice1)
	fmt.Println("修改slicen后，数组是", arr1)
	//可以查看一下数组中元素的地址和对应的切片中的元素地址，进行对比.可以看到地址是相同的
	fmt.Printf("arr1[1]的地址是%p\n", &arr1[1])
	fmt.Printf("slice1[0]的地址是%p\n", &slice1[0])
}
