package unbuffered

import (
	"fmt"
	"sync"
)

type UnbufferedFooBar struct {
	fooChan chan int
	barChan chan int
	wg      *sync.WaitGroup
}

func NewChannel(wg *sync.WaitGroup) *UnbufferedFooBar {
	return &UnbufferedFooBar{
		fooChan: make(chan int),
		barChan: make(chan int),
		wg:      wg,
	}
}

func (this *UnbufferedFooBar) KickStart() {
	this.fooChan <- 1
}

func (this *UnbufferedFooBar) Foo(n int) {
	defer this.wg.Done()
	for i := 0; i < n; i++ {
		<-this.fooChan
		fmt.Print("foo")
		this.barChan <- 1
	}
}

func (this *UnbufferedFooBar) Bar(n int) {
	defer this.wg.Done()
	for i := 0; i < n; i++ {
		<-this.barChan
		fmt.Print("bar")

		if i < n-1 {
			this.fooChan <- 1
		}
	}
}
