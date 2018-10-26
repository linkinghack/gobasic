package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal // 操作系统发送的信号

	complete chan error // 用于报告处理任务已完成

	timeout <-chan time.Time //报告任务已超时

	tasks []func(int) // 任务函数
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

// New 工厂方法，返回一个新的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), // 缓冲大小为1， 至少可接受一个操作系统发来的interrupt
		complete:  make(chan error),        // 无缓冲通道
		timeout:   time.After(d),           // After() 返回一个 <-chan time.Time， 当d时间过去后
		// task 字段默认的nil
	}
}

// Add 将任意多个func(int) 作为任务添加到runner对象中
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...) // 注意append() 方法将创建新对象
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt) // 接收中断信号, 系统Interrupt 只想自定义Interrupt （os.Interrupt接口）

	// 用不同goroutine处理任务
	go func() {
		r.complete <- r.run() //run() 的运行过程不会阻塞当前方法Start()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测系统中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//若没同时中断则执行任务
		task(id)
	}
	return nil // 无错误，未被中断
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}
