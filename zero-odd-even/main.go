package main

import (
	"fmt"
	"sync"
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
}
