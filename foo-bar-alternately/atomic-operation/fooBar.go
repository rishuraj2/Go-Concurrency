package atomicoperation

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicFooBar struct {
	n       int
	wg      *sync.WaitGroup
	fooTurn atomic.Bool
}

func NewAtomicFooBar(wg *sync.WaitGroup, n int) *AtomicFooBar {
	fb := &AtomicFooBar{
		wg: wg,
		n:  n,
	}

	fb.fooTurn.Store(true)
	return fb
}

func (this *AtomicFooBar) Foo() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		for !this.fooTurn.Load() {
			// I am waiting and burning cpu :)
		}
		fmt.Print("foo")
		this.fooTurn.Store(false) // Hey Bar! take over the control
	}
}

func (this *AtomicFooBar) Bar() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		for this.fooTurn.Load() {
			// I am waiting and burning cpu :)
		}
		fmt.Print("bar")
		this.fooTurn.Store(true) // Hey Foo! it's your turn now
	}
}
