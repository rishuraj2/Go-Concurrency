package buffered

import (
	"fmt"
	"sync"
)

type BufferedFooBar struct {
	fooChan chan int
	barChan chan int
	n       int
	wg      *sync.WaitGroup
}

func NewBufferedFooBar(wg *sync.WaitGroup, n int) *BufferedFooBar {
	fb := &BufferedFooBar{
		fooChan: make(chan int, 1),
		barChan: make(chan int, 1),
		n:       n,
		wg:      wg,
	}

	fb.fooChan <- 1 // ensuring that foo goes first

	return fb
}

func (this *BufferedFooBar) Foo() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		<-this.fooChan
		fmt.Print("foo")
		this.barChan <- 1
	}
}

func (this *BufferedFooBar) Bar() {
	defer this.wg.Done()
	for i := 0; i < this.n; i++ {
		<-this.barChan
		fmt.Print("bar")
		this.fooChan <- 1
	}
}
