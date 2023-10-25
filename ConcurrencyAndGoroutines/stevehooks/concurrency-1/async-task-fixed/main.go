package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := make(chan struct{})
	/* go func() {
		task1(ch)
		<-ch
	}()
	*/
	go task1(ch)
	go task2(ch)
	go task3(ch)
	go task4(ch)
	<-ch
	<-ch
	<-ch
	<-ch
	//time.Sleep(time.Second)
	fmt.Println("Elapsed", time.Since(now))

}

func task1(ch chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
	ch <- struct{}{}
}
func task2(ch chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2")
	ch <- struct{}{}

}
func task3(ch chan struct{}) {
	fmt.Println("task3")
	ch <- struct{}{}
}
func task4(ch chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
	ch <- struct{}{}
}
