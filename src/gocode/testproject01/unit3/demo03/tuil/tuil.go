package tuil

import "fmt"

var (
	Name string
	Age  int
	Sex  string
)

func init() {
	fmt.Println("这是util包中的init函数")
	Age = 28
	Name = "zhaoyan"
	Sex = "男"

}
