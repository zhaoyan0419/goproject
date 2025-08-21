package main

import (
	"fmt"
	"strconv"
	"strings"
)

func test() {
	// string转int
	var str1 string = "hello world你好"
	fmt.Println(len(str1))
	for k, v := range str1 {
		fmt.Printf("str1[%d]:%c     \n", k, v)
	}
}
func str2int() {
	str1 := "666"
	num1, _ := strconv.Atoi(str1)
	fmt.Println(num1)
}

func int2str() {
	//int转string
	str := strconv.Itoa(123)
	fmt.Println(str)
}

func countstr() {

	// 返回一个字符串里包含另一个字符串的个数
	count := strings.Count("hello golang,where are you go go gogogo", "go")
	fmt.Println(count)
}

func strstr() {
	//展示不区分大小写的字符串比较
	equal := strings.EqualFold("Go", "go")
	fmt.Println(equal)
}

func strstr2() {
	//区分大小写的字符串比较
	fmt.Println("hello" == "HELLO")
}

func strFirstIndex() {
	//返回一个子字符串在字符串中第一次出现的索引/如果返回-1，则子串不存在
	indexs := strings.Index("gfsdagsdfasdfstrbghdfgh", "str")
	fmt.Println(indexs)
}

func strReplace() {
	// 字符串替换，参数最后的int是替换字符串的个数，-1代表全部替换
	str := strings.Replace("strjkljknjstrljl;n;lstrljlkstrljkljlnstrnlk;lstrl;l;kjlstr", "str", "go", 3)
	fmt.Println(str)
}
func strAwk() {
	//类似于awk -F “ ” 将字符串按照指定分隔符进行分割并且存入一个数组
	arr := strings.Split("afldsskj lksjfdljdskl lkjlkfdl lkjljlj llk l l o o", " ")
	fmt.Println(arr)
}

func LowerOrUpper() {
	// 对字符串进行大小写转换
	str1 := strings.ToUpper("StringGolang")
	str2 := strings.ToLower("StringGolang")
	fmt.Println("我是将字符串全部变成大写的函数", str1)
	fmt.Println("我是将字符串全部变成小写的函数", str2)
}

func rmRightAndLeftSpaceInString() {
	//去掉字符串左右的空格
	str := strings.TrimSpace("     go  java  python     ")
	println(str)
}
func rmRightAndLeftStrInStr() {
	//去掉字符串左右的指定字符或字符串
	str := strings.Trim("go java go java go", "go")
	fmt.Println(str)
}
func rmLeftStrInStr() {
	//去除左边的指定字符
	str := strings.TrimLeft("go java go", "go")
	fmt.Println(str)
}
func rmRightStrInStr() {
	//去除右边的指定字符
	str := strings.TrimRight("go java go", "go")
	fmt.Println(str)
}
func beginAtStr() {
	// 判断一个字符串是否以xxx开头
	trueOrFalse := strings.HasPrefix("http://tengxunyun.zhaoyan.site", "http")
	fmt.Println(trueOrFalse)
}

func endAtStr() {
	//判断某个字符串是否以xxx结尾
	trueOrFalse := strings.HasSuffix("demo.jpg", "jpg")
	fmt.Println(trueOrFalse)
}
func main() {
	beginAtStr()
	endAtStr()
}
