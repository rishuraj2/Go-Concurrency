package sharedstate

import (
	"fmt"
	"sync"
)

type Signaling struct {
	mu        sync.Mutex
	cond      *sync.Cond
	wg        *sync.WaitGroup
	isFooTurn bool
}

func NewSignaling(wg *sync.WaitGroup) *Signaling {
	sig := &Signaling{
		isFooTurn: true,
		wg:        wg,
	}

	sig.cond = sync.NewCond(&sig.mu)
	return sig
}

func (this *Signaling) Foo(n int) {
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

func (this *Signaling) Bar(n int) {
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
