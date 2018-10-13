package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time" // time.sleep()阻塞进程
)

var (
	// shutdown 用于通知正在执行的goroutine停止工作
	shutdown int64

	// 等待所有goroutine结束
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("Tom")
	go doWork("Jerry")

	// 等待所有goroutine1秒
	time.Sleep(1 * time.Second)

	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
	fmt.Println("所有Goroutine结束")
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work \n", name)
		time.Sleep(300 * time.Millisecond) //阻塞300毫秒

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shuting down %s work\n", name)
			break
		}
	}
}
