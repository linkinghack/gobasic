package main

/**
* 缓冲通道模拟接力跑,利用递归
**/
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1 //初始化第一位队员

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton //等待接力棒     ------（1）

	fmt.Printf("Runner %d running with baton\n", runner)

	// 创建下一位跑步队员
	// 采用递归的方式创建四个goroutine，每个代表一个接力队员
	// 创建的另外1个goroutine将在 (1) 的位置停下等待，第二个goroutine再次运行至此
	// 第一位已timeSleep完毕，此时才创建第三个goroutine
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// 跑步
	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done() // runner4结束时才计数减1
		return
	}

	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)

	baton <- newRunner
}
