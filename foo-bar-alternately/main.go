package main

import (
	sharedstate "foobar/shared-state"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sig := sharedstate.NewSignaling(&wg)
	n := 10

	wg.Add(2)
	go sig.Foo(n)
	go sig.Bar(n)
	wg.Wait()
}