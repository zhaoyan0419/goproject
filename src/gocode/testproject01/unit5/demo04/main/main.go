package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	arr2 := [3]int{9, 8, 7}
	fmt.Println(arr2)
	var arr3 = [...]int{52, 4, 8, 3, 4}
	fmt.Println(arr3)
	var arr4 = [...]int{2: 66, 0: 33, 1: 789, 3: 987}
	fmt.Println(arr4)
}
