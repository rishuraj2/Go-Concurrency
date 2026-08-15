package unbuffered

import (
	"fmt"
	"sync"
)

type UnbufferedFooBar struct {
	fooChan chan int
	barChan chan int
	n       int
	wg      *sync.WaitGroup
}

func NewUnbufferedFooBar(wg *sync.WaitGroup, n int) *UnbufferedFooBar {
	return &UnbufferedFooBar{
		fooChan: make(chan int),
		barChan: make(chan int),
		n:       n,
		wg:      wg,
	}
}

func (this *UnbufferedFooBar) KickStart() {
	this.fooChan <- 1
}

func (this *UnbufferedFooBar) Foo() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		<-this.fooChan
		fmt.Print("foo")
		this.barChan <- 1
	}
}

func (this *UnbufferedFooBar) Bar() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		<-this.barChan
		fmt.Print("bar")

		if i < this.n-1 {
			this.fooChan <- 1
		}
	}
}
