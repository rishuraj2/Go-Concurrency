package sharedmemory

import (
	"fmt"
	"sync"
)

type SharedMemFooBar struct {
	mu        sync.Mutex
	cond      *sync.Cond
	wg        *sync.WaitGroup
	isFooTurn bool
}

func NewSignaling(wg *sync.WaitGroup) *SharedMemFooBar {
	sig := &SharedMemFooBar{
		isFooTurn: true,
		wg:        wg,
	}

	sig.cond = sync.NewCond(&sig.mu)
	return sig
}

func (this *SharedMemFooBar) Foo(n int) {
	defer this.wg.Done()

	for i := 0; i < n; i++ {
		this.mu.Lock()

		for !this.isFooTurn {
			this.cond.Wait()
		}
		fmt.Print("foo")

		this.isFooTurn = false
		this.cond.Signal()
		this.mu.Unlock()
	}
}

func (this *SharedMemFooBar) Bar(n int) {
	defer this.wg.Done()

	for i := 0; i < n; i++ {
		this.mu.Lock()

		for this.isFooTurn {
			this.cond.Wait()
		}

		fmt.Print("bar")
		this.isFooTurn = true
		this.cond.Signal()
		this.mu.Unlock()
	}
}
