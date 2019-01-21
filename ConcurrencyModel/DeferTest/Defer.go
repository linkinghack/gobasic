package main

import (
	"fmt"
	"time"
)

func longTimeConsumeTask() {
	fmt.Println("long time consume task doing>>..")
	time.Sleep(time.Second * 5)
}

func start() string {
	defer longTimeConsumeTask()
	fmt.Println("program start")
	time.Sleep(time.Second * 3)
	return "returned"
}

func main() {
	fmt.Println(start())
}
