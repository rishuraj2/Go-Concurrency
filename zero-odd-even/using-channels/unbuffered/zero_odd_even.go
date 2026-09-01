package unbuffered

import (
	"fmt"
	"sync"
)

type ZeroOddEven struct {
	n        int
	evenChan chan int
	oddChan  chan int
	zeroChan chan int
	wg       *sync.WaitGroup
}

func NewZeroOddEven(n int, wg *sync.WaitGroup) *ZeroOddEven {
	return &ZeroOddEven{
		n:        n,
		evenChan: make(chan int),
		oddChan:  make(chan int),
		zeroChan: make(chan int),
		wg:       wg,
	}
}

func (this *ZeroOddEven) Zero() {
	defer this.wg.Done()

	for i := 1; i <= this.n; i++ {
		fmt.Print("0")
		if i%2 == 0 {
			this.evenChan <- i
		} else {
			this.oddChan <- i
		}

		<-this.zeroChan
	}

	close(this.evenChan)
	close(this.oddChan)
	close(this.zeroChan)
}

func (this *ZeroOddEven) Odd() {
	defer this.wg.Done()
	for n := range this.oddChan {
		fmt.Print(n)
		this.zeroChan <- 0
	}
}

func (this *ZeroOddEven) Even() {
	defer this.wg.Done()
	for n := range this.evenChan {
		fmt.Print(n)
		this.zeroChan <- 0
	}
}
