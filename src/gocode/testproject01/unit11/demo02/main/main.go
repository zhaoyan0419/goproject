package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ { // 匿名函数+外部变量 = // 闭包
		wg.Add(1) // 协程开始时候加1操作
		// 启动一个协程
		// 使用匿名函数，直接调用个匿名函数
		go func() {
			defer wg.Done()
			fmt.Println(i)
			//wg.Done() //协程执行完成减1
		}()
		//time.Sleep(time.Second)
	}

	wg.Wait()
	// 等待一秒是因为主死从随，给主函数阻塞一秒，用于给协程将程序执行完毕
	// 但是在开发过程中，通常不知道协程究竟需要执行多久，主进程需要等待多久
	//time.Sleep(time.Second * 1)

}
