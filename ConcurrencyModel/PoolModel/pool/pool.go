package pool

/**
* Pool is a goroutine pool with buffered channel
**/

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer // 资源要实现Closer接口，包含Close()方法，即自己把自己关闭
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New 创建一个资源池
// 第一个参数是一个可以用来分配新资源的函数
// 第二个参数规定池大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size cannot be negative or zero.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resources")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil // 正常时error为nil即可，没有消息就是好消息

	// 没有空闲资源可用，提供一个新资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory() // factory() 时在New资源池时指定的资源工厂
	}
}

// 将一个使用完的资源回收，
// 在池已经被关闭的情况下，直接关闭该资源
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close()操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 池已经被关闭，则关闭资源
	if p.closed {
		r.Close()
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	default:
		log.Println("Release:", "Closing") // 参考下方Close()方法
	}
}

// 关闭资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return // 已关闭
	}

	//置状态
	p.closed = true

	// 清空通道之前应该关闭通道
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
