package main

/**
* 经典“生产者-消费者”问题， golang 有缓冲通道方式
**/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberOfGoroutines = 4  //goroutine 数量(消费者)
	taskLoad           = 10 // 任务数量(生产者)
)

var wg sync.WaitGroup // goroutine 数量信号量，阻塞主进程

// Main之前执行
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建缓冲通道来管理任务
	tasks := make(chan string, taskLoad)

	wg.Add(numberOfGoroutines)
	// 启动goroutines来处理任务
	for gid := 1; gid <= numberOfGoroutines; gid++ {
		go worker(tasks, gid)
	}

	// 添加任务
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	close(tasks)

	wg.Wait()
	fmt.Println("All done")
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// 通道关闭，没有多余任务
			fmt.Printf("Worker: %d Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d Started %s\n", worker, task)

		// 随机等待一段时间来模拟工作
		sleepTime := rand.Int63n(100)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

		// 输出完成结果
		fmt.Printf("Worker: %d Completed %s\n", worker, task)
	}
}
