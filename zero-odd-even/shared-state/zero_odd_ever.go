package sharedstate

import (
	"fmt"
	"sync"
)

type ZeroOddEven struct {
	n     int
	state State
	mu    sync.Mutex
	wg    *sync.WaitGroup
	cond  *sync.Cond
}

func NewZeroOddEven(n int, wg *sync.WaitGroup) *ZeroOddEven {
	instance := &ZeroOddEven{
		n:     n,
		state: ZERO_BEFORE_ODD,
		wg:    wg,
	}

	instance.cond = sync.NewCond(&instance.mu)
	return instance
}

func (this *ZeroOddEven) Zero() {
	defer this.wg.Done()

	for i := 0; i < this.n; i++ {
		this.mu.Lock()

		for this.state != ZERO_BEFORE_EVEN && this.state != ZERO_BEFORE_ODD {
			this.cond.Wait()
		}

		fmt.Print("0")

		if this.state == ZERO_BEFORE_EVEN {
			this.state = EVEN
		} else {
			this.state = ODD
		}

		this.cond.Broadcast()
		this.mu.Unlock()
	}
}

func (this *ZeroOddEven) Odd() {
	defer this.wg.Done()
	current := 1

	for current <= this.n {
		this.mu.Lock()

		for this.state != ODD {
			this.cond.Wait()
		}

		fmt.Print(current)
		current += 2
		this.state = ZERO_BEFORE_EVEN

		this.cond.Broadcast()
		this.mu.Unlock()
	}

}

func (this *ZeroOddEven) Even() {
	defer this.wg.Done()
	current := 2

	for current <= this.n {
		this.mu.Lock()
		for this.state != EVEN {
			this.cond.Wait()
		}

		fmt.Print(current)
		current += 2
		this.state = ZERO_BEFORE_ODD
		this.cond.Broadcast()
		this.mu.Unlock()
	}

}
