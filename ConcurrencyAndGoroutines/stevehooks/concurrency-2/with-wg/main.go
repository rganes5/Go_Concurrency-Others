package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	fmt.Println("number of cores", runtime.NumCPU())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(i+1, &wg)
	}
	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("Main is done, elapsed", time.Since(now))
}

func work(id int, wg *sync.WaitGroup) {
	//time.Sleep(100 * time.Millisecond)
	defer wg.Done()
	fmt.Println("task is", id, "done")
}
