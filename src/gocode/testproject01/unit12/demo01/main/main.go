package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	// 对基本数据类型进行反射
	var num1 int = 100
	tp := reflect.TypeOf(num1)
	fmt.Println(tp)
	value := reflect.ValueOf(num1)
	fmt.Println(value)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	tp1 := reflect.TypeOf(reader)
	fmt.Println(tp1)
	value1 := reflect.ValueOf(reader)
	fmt.Println(value1)
	fmt.Println(str)
}
