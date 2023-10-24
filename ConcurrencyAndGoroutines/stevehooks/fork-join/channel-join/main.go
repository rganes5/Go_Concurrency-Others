package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	now := time.Now()
	//fork point
	go func() {
		work()
		ch <- struct{}{}
	}()
	//join point
	<-ch
	fmt.Println("elapsed", time.Since(now))

	fmt.Println("done waiting, main exits")
}

func work() {
	fmt.Println("from the work")
}
