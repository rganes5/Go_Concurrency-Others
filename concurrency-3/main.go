package main

import (
	"fmt"
	"sync"
)

var counter int64

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("before", counter)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			Inc(&mu)
		}(&wg)
	}
	wg.Wait()
	fmt.Println("after", counter)

}

func Inc(mu *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
