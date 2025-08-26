package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()

		totalNum = totalNum + 1
		mutex.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		totalNum = totalNum - 1
		mutex.Unlock()

	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()

	wg.Wait()
	// 此处输出的结果不一定是0，因为两个协程是同步执行的，有可能两个协程最开始拿到的totalnum是相同的，比如都是1
	// 那么如果sum先执行完之后给totalnum赋值为2
	// 但是sub后执行完，sub拿到的初始值是1，-1之后是0，sub会将0赋值给totalnum，会将sum函数赋给totalnum的值覆盖
	// 所以最终输出的结果不一定是0
	fmt.Println(totalNum)
}
