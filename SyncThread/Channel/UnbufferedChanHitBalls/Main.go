package main

/**
* 使用无缓冲channel模拟打网球
**/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() { // init方法将会在main之前执行
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int) // 无缓冲通道

	wg.Add(2)

	go player("二货王小萌", court)
	go player("聪明的刘磊", court)

	// 在此之前court是空的，两个goroutine都阻塞
	court <- 1 //发球，

	wg.Wait()
	fmt.Println("游戏结束")

}

func player(name string, court chan int) {
	defer wg.Done()

	for { //在此游戏过程中仅有两个出口，要么输要么赢，否则出手并继续

		ball, ok := <-court
		if !ok { // 通道关闭则赢得比赛
			fmt.Printf("Player %s Win\n", name)
			return
		}

		n := rand.Intn(100) //0-100随机数
		if n%13 == 0 {
			fmt.Printf("Player %s Missed(get %d) \n", name, n)

			close(court) // 关闭通道，输掉了
			return
		}

		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		court <- ball // 将球打给对手
	}
}
