package main

import (
	"errors"
	"fmt"
)

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数是0，请检查除数值")
	} else {
		result := num1 / num2
		fmt.Println(result)

		return nil
	}

}

// 使用panic函数之后，程序不继续向下走
func main() {
	err := test()
	if err != nil {
		fmt.Println(error(err))
		panic(err)
	}
	fmt.Println("上边的除法操作成功，正常执行下边的逻辑")
}
