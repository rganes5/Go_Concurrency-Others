package main

import (
	"fmt"
	"time"
)

func main() {
	//	ch := make(chan struct{})
	//go work(ch)
	//ch <- struct{}{}
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting")
}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("from the work")
	//<-ch

}
