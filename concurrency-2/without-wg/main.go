package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("number of cores", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main is done, elapsed", time.Since(now))
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task is", id, "done")
}
