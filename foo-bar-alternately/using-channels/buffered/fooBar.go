package buffered

import (
	"fmt"
	"sync"
)

type BufferedFooBar struct {
	fooChan chan int
	barChan chan int
	wg      *sync.WaitGroup
}

func NewChannel(wg *sync.WaitGroup) *BufferedFooBar {
	return &BufferedFooBar{
		fooChan: make(chan int, 1),
		barChan: make(chan int, 1),
		wg:      wg,
	}
}

func (this *BufferedFooBar) KickStart() {
	this.fooChan <- 1
}

func (this *BufferedFooBar) Foo(n int) {
	defer this.wg.Done()
	for i := 0; i < n; i++ {
		<-this.fooChan
		fmt.Print("foo")
		this.barChan <- 1
	}
}

func (this *BufferedFooBar) Bar(n int) {
	defer this.wg.Done()
	for i := 0; i < n; i++ {
		<-this.barChan
		fmt.Print("bar")
		this.fooChan <- 1
	}
}
