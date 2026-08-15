package main

import (
	"fmt"
	sharedmemory "foobar/shared-memory"
	"foobar/using-channels/buffered"
	"foobar/using-channels/unbuffered"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := 10

	fmt.Println("---------- Shared State ----------")
	sig := sharedmemory.NewSignaling(&wg)
	wg.Add(2)
	go sig.Foo(n)
	go sig.Bar(n)
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Message Passing [buffered channel] ----------")
	bufChan := buffered.NewChannel(&wg)
	wg.Add(2)
	go bufChan.Foo(n)
	go bufChan.Bar(n)
	bufChan.KickStart()
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Message Passing [unbuffered channel] ----------")
	unbufChan := unbuffered.NewChannel(&wg)
	wg.Add(2)
	go unbufChan.Foo(n)
	go unbufChan.Bar(n)
	unbufChan.KickStart()
	wg.Wait()
	fmt.Print("\n\n")
}
