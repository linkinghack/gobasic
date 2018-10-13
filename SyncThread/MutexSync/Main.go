package main

/**
Note:  Mutex作为一种临界资源锁，和具体资源并不绑定，它只关注过程；
	一个mutex对象就是一个独立的锁，它的状态在不同的方法创建的goroutine中都有效；
	所以可以实现生产者消费者问题的复现
**/

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int //临界资源
	wg      sync.WaitGroup
	mutex   sync.Mutex //准备一个锁
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("final counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 创建一个临界区，同一时刻只允许一个goroutine进入
		mutex.Lock()
		{
			value := counter

			runtime.Gosched() // 当前线程退出到就绪队列

			value++

			counter = value
		}
		mutex.Unlock() // 解开临界锁
	}
}
