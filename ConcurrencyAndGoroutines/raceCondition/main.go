package main

import (
	"fmt"
	"sync"
)

// race condition example
func main() {

	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	score := []int{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine :1")
		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine :2")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()

	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Go routine :3")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()

	}(wg, mu)

	wg.Wait()
	fmt.Println("final slice:", score)

}
