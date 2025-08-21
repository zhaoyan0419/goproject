package main

import "fmt"

// map的引入

func testMap1() {
	var map1 map[int]string
	// map在使用前必须要先make，make的第二个参数可以省略，如果省略，默认会分配一个内存
	map1 = make(map[int]string, 10)
	map1[201634] = "zhaoyan"
	map1[201633] = "hzy"
	map1[201632] = "ts"
	// 重复对一个key进行存储，会进行替换
	map1[201634] = "zs"
	fmt.Println(map1)
	fmt.Println(len(map1))
	// 方式二，使用类型推断定义一个map
	map2 := make(map[int]string, 10)
	fmt.Println(map2)
	// 方式三，用大括号
	map3 := map[int]string{
		123321: "zhaoyan",
		111222: "wangwu",
		321321: "lisi",
	}
	fmt.Println(map3)

	// 超长度测试
	testMap := make(map[int]string, 1)
	testMap[1] = "zzz"
	testMap[2] = "sss"
	fmt.Println(testMap)
}

// map的增删改查
func mapCURD() {
	testMap := make(map[int]string, 10)
	// 增
	testMap[1] = "zzz"
	testMap[2] = "sss"
	fmt.Println(testMap)
	//改
	testMap[2] = "test"
	fmt.Println(testMap)
	// 删
	delete(testMap, 1)
	fmt.Println(testMap)
	// 查
	value, flag := testMap[2]
	fmt.Println(value)
	fmt.Println(flag)
}
func main() {
	// map遍历
	testMap := make(map[int]string)
	for i := 0; i < 10; i++ {
		testMap[i] = fmt.Sprintf("the num of string is %v", i)

	}
	fmt.Println(testMap)
	fmt.Println(len(testMap))
	for k, v := range testMap {
		fmt.Printf("testMap[%v] = %s\n", k, v)
	}
	// 加深难度
	a := make(map[string]map[int]string)
	b := make(map[int]string)
	b[100] = "zy"
	b[200] = "zd"
	b[300] = "wqs"
	c := make(map[int]string)
	c[100] = "hzy"
	c[200] = "ts"
	c[300] = "fq"
	a["班级1"] = b
	a["班级2"] = c
	fmt.Println(a)
	for i, j := range a {
		fmt.Printf("以下遍历%s\n", i)
		for k, v := range j {
			fmt.Printf("%d = %s\n", k, v)
		}
	}
}
