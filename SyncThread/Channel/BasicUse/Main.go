package main

import (
	"fmt"
	"time"
)

func tickBoom() {
	tick := time.Tick(30 * time.Millisecond)
	boom := time.After(300 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("...")
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	c <- 3
	fmt.Println(<-c)
	close(c)
	c <- 1
}
