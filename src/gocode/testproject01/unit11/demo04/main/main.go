package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second * 3)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}

func main() {
	wg.Add(5)
	read()
	go read()
	go read()
	go read()
	go read()
	go write()
	wg.Wait()
}
