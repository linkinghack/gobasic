package work

/**
* work is a goroutine pool model with unbuffered channel
**/
import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// Although the channel is closed, this may not be terminated immediately, will not cause a panic
			for w := range p.work { // this will consumes all the workers in the channel(p.work),
				w.Task()
			}
			// Every goroutine are reading works fromt the channel, and Task() them concurrently
			p.wg.Done() // this goroutine worked done
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work) // close the channel in case of the memory leak
	p.wg.Wait()   // Wait until all the goroutines consumes all work remained in the channel have done their jobs
}
