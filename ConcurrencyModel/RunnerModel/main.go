package main

import (
	"log"
	"os"
	"time"

	"./runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// 初始化本次任务并分配超时时间
	r := runner.New(timeout)

	// 加入任务
	r.Add(createTask(), createTask(), createTask())

	//处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
}

func createTask() func(int) { //返回值是函数类型
	return func(id int) {
		log.Println("Processor - task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}

}
