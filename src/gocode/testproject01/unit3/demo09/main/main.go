package main

import "fmt"

func testLenth() {
	str := "go java let up study golang go语言"
	fmt.Println(len(str))
}
func testNew() {
	num := new(int)
	fmt.Printf("num 的值为%v，num的类型为%T，num的地址为%v，num存的地址内的值为%v\n", num, num, &num, *num)

}
func main() {
	testNew()
}
