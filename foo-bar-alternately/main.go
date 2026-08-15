package main

import (
	"fmt"
	atomicoperation "foobar/atomic-operation"
	sharedmemory "foobar/shared-memory"
	"foobar/using-channels/buffered"
	"foobar/using-channels/unbuffered"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := 10

	fmt.Println("---------- Atomic flag ----------")
	at := atomicoperation.NewAtomicFooBar(&wg, n)
	wg.Add(2)
	go at.Foo()
	go at.Bar()
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Message Passing [buffered channel] ----------")
	bufChan := buffered.NewBufferedFooBar(&wg, n)
	wg.Add(2)
	go bufChan.Foo()
	go bufChan.Bar()
	bufChan.KickStart()
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Message Passing [unbuffered channel] ----------")
	unbufChan := unbuffered.NewUnbufferedFooBar(&wg, n)
	wg.Add(2)
	go unbufChan.Foo()
	go unbufChan.Bar()
	unbufChan.KickStart()
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Shared State ----------")
	sig := sharedmemory.NewSharedMemoryFooBar(&wg, n)
	wg.Add(2)
	go sig.Foo()
	go sig.Bar()
	wg.Wait()
	fmt.Print("\n\n")
}
