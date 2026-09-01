package main

import (
	"fmt"
	"sync"
	sharedstate "zerooddeven/shared-state"
	"zerooddeven/using-channels/unbuffered"
)

func main() {
	var wg sync.WaitGroup
	n := 5

	fmt.Println("---------- Message Passing [unbuffered channel] ----------")
	unbufChan := unbuffered.NewZeroOddEven(n, &wg)
	wg.Add(3)
	go unbufChan.Even()
	go unbufChan.Odd()
	go unbufChan.Zero()
	wg.Wait()
	fmt.Print("\n\n")

	fmt.Println("---------- Shared State ----------")
	shareState := sharedstate.NewZeroOddEven(n, &wg)
	wg.Add(3)
	go shareState.Odd()
	go shareState.Even()
	go shareState.Zero()
	wg.Wait()
	fmt.Print("\n\n")
}
