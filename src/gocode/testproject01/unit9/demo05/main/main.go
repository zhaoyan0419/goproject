package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("Name = %v, Age = %v", s.Name, s.Age)
	return str
}

// 如果一个类型实现了String() string 这个方法,那么fmt.println默认会调用这个变量的String()进行输出
// 以后定义结构体的化，常定义String()这个方法作为输出结构体信息的方法，在fmt.println时候会自动调用
func main() {
	s1 := Student{
		Name: "zhaoyan",
		Age:  28,
	}
	// 传入地址，如果绑定了String方法就会自动调用
	fmt.Println(&s1)

}
