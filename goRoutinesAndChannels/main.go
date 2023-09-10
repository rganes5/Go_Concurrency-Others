package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//Unidirectional channel
	/*
		r:= make(<-chan int) //read only
		s:= make(chan <-int) //write only
	*/

	//Simple Demo (Waitgroup not necessarily needed)
	/*
		cha := make(chan int)
		//wg := sync.WaitGroup{}

		//wg.Add(1)
		go func() {
			//	defer wg.Done()
			cha <- 5 //Sending data to a channel / Write to the channel

		}()
		value := <-cha //Receving data from a channel / Read from a channel
		fmt.Println(value)

		//wg.Wait()
	*/

	sendChan := make(chan int)
	go sendVal(sendChan)
	for i := 0; i < 6; i++ {
		fmt.Println(<-sendChan)
	}

}

func sendVal(sendChan chan int) {

	for i := 0; i < 5; i++ {
		sendChan <- i
	}
	close(sendChan)
}
