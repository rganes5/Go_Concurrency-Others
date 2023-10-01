package main

import (
	"fmt"
	"sync"
)

// waitgroup
func main() {
	slice := make([]int, 100)
	fmt.Println(slice)

	var wg sync.WaitGroup
	fchann := make(chan int)
	rchann := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fchann <- i + 1
		}
		close(fchann)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range fchann {
			rchann <- val
		}
		close(rchann)
	}()

	for i := 0; i < 100; i++ {
		slice[i] = <-rchann
	}
	fmt.Println(slice)
	wg.Wait()

}
