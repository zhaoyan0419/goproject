package main

import "fmt"

func student1() {
	fmt.Println("请输入学生姓名：")
	var name string
	fmt.Scanln(&name)

	fmt.Println("请输入学生成绩：")
	var score float32
	fmt.Scanln(&score)

	fmt.Println("请输入学生是否为vip true/false：")
	var vip bool
	fmt.Scanln(&vip)

	fmt.Println("请输入学生年龄：")
	var age int
	fmt.Scanln(&age)

	fmt.Printf("学生的姓名为:%v，学生的年龄为:%v，学生的成绩为:%v，学生是否为vip:%v", name, age, score, vip)
}

func student2() {
	var (
		name  string
		age   int
		score float32
		vip   bool
	)
	fmt.Println("请输入学生的姓名 年龄 分数 和是否为vip，用空格分隔")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &vip)
	fmt.Printf("学生的姓名为:%v，学生的年龄为:%v，学生的成绩为:%v，学生是否为vip:%v", name, age, score, vip)
}

func main() {
	student2()
}
