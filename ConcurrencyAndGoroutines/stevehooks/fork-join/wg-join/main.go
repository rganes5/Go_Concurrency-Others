package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()
	fmt.Println("elapsed", time.Since(now))

	fmt.Println("done waiting, main exits")
}

func work() {
	fmt.Println("from the work")
}
