package main

import (
	"fmt"
	"sync"
	"time"
)

var i int
var wg sync.WaitGroup
var clo sync.Mutex

func add() {
	//defer wg.Done()
	clo.Lock()
	i = i + 1
	time.Sleep(time.Second * 10)
	clo.Unlock()
}
func sub() {
	defer wg.Done()
	clo.Lock()
	i = i + 100
	fmt.Println(i)
	clo.Unlock()
}

func main() {
	wg.Add(1)

	add()
	// 在add函数中加了互斥锁，在执行add函数执行的过程中，下边的协程需要等待，需要等待add函数释放了互斥锁
	go sub()
	wg.Wait()
}
