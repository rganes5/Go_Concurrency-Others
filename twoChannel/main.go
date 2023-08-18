package main

import (
	"fmt"
)

// Two channel communication
func main() {
	fmt.Println("Demo of two channel in golang")

	sCh := make(chan int)
	rCh := make(chan int)
	slice := make([]int, 100)
	fmt.Println(slice)

	go func() {
		for i := 0; i < 100; i++ {
			sCh <- i + 1
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			value := <-sCh
			rCh <- value
		}
	}()

	for i := 0; i < 100; i++ {
		slice[i] = <-rCh
	}

	fmt.Println(slice)
}
