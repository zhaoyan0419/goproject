package main

import "fmt"

type Teacher struct {
	name   string
	age    int
	school string
}

func makeStruct() {
	var t1 Teacher = Teacher{
		name:   "zhaoyan",
		age:    28,
		school: "sie",
	}
	t1.name = "zhangsan"
	fmt.Println(t1)
	// 默认值
	var t2 Teacher
	fmt.Println(t2)
	// 方式二
	t2.name = "hzy"
	fmt.Println(t2.name)
	// 方式3
	var t3 *Teacher = new(Teacher)
	// t3是指针，指向的是地址，应该给这个地址指向的对象字段进行赋值
	fmt.Println(t3)
	(*t3).age = 45
	// 为了符合程序员的编程习惯，go提供了简化的赋值方式
	t3.name = "mashibing"
	fmt.Println(t3)
	fmt.Println(&t3)
	// 方式4
	var t4 *Teacher = &Teacher{name: "asdf", age: 32, school: "sdfdsf"}
	fmt.Printf("t4 : %v\n", *t4)

	//t4.name = "wangqiushi"
	fmt.Println(t4)
}

func testStructPtr() {
	a := new(int)
	b := new(Teacher)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("the ptr of b is %v\n", *b)
}

func main() {
	testStructPtr()
}
