package main

import (
	"fmt"
	"time"
)

func nowTime() {
	//获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("now 的类型是：%T\n", now)
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Printf("月的类型是：%T\n", now.Month())
	fmt.Println("月：", int(now.Month()))
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

}

func dateFmt() {
	// 日期的格式化，按照自己预期的方式输出时间为一个字符串
	fmt.Printf("获取当前年月日%d-%d-%d\n", int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day()))
	dataStr := fmt.Sprintf("获取当前年月日%d-%d-%d", int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day()))
	fmt.Println(dataStr)
}

func dateFmt1() {
	nowtime := time.Now().Format("20060102150405")
	fmt.Println(nowtime)
}

func dateFmt3() {
	nowtime := time.Now().Format("2006:01:02")
	fmt.Println(nowtime)
}

func main() {

}
