package main

import "fmt"

// if单分支
func test() {

	if count := 25; count < 30 {
		fmt.Println("口罩库存不足")
	}
}

// if双分支
func test2() {
	var count int
	count = 31
	if count < 30 {
		fmt.Println("库存不足")
	} else {
		fmt.Printf("库存充足，口罩剩余%d\n", count)
	}
}

// if多分支
func test3() {

	if score := 12; score >= 90 {
		fmt.Println("成绩为优秀")
	} else if score >= 60 {
		fmt.Println("成绩为及格")
	} else {
		fmt.Println("成绩为不及格")
	}
}

func test4() {
	var score int32 = 111
	var b int64 = 9
	switch score / 10 {
	case 10:
		fmt.Println("等级为A")
	case int32(b):
		fmt.Println("等级为B")
	case 8:
		fmt.Println("等级为C")
	case 7:
		fmt.Println("等级为D")
	case 6:
		fmt.Println("等级为E")
	case 5:
		fmt.Println("等级为F")
	case 4:
		fmt.Println("等级为G")
	case 3:
		fmt.Println("等级为H")
	case 2:
		fmt.Println("等级为I")
	case 1:
		fmt.Println("等级为J")
	case 0:
		fmt.Println("等级为K")
	default:
		fmt.Println("您的成绩有误")
	}

}
func test5() {
	switch a := 1; {
	case a == 1:
		fmt.Println("a=1")
		fallthrough
	default:
		fmt.Println("error")
	}
}
func main() {
	test5()
}
