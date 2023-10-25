package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait()
}
