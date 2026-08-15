package sharedmemory

import (
	"fmt"
	"sync"
)

type SharedMemFooBar struct {
	mu        sync.Mutex
	cond      *sync.Cond
	n         int
	wg        *sync.WaitGroup
	isFooTurn bool
}

func NewSharedMemoryFooBar(wg *sync.WaitGroup, n int) *SharedMemFooBar {
	sig := &SharedMemFooBar{
		isFooTurn: true,
		n:         n,
		wg:        wg,
	}

	sig.cond = sync.NewCond(&sig.mu)
	return sig
}

func (this *SharedMemFooBar) Foo() {
	defer this.wg.Done()

	for i := 0; i < this.n; i++ {
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

func (this *SharedMemFooBar) Bar() {
	defer this.wg.Done()

	for i := 0; i < this.n; i++ {
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
