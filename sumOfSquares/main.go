// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	number := make(chan int)
	response := make(chan int)
	slice := []int{1, 2, 3, 4, 5}
	total := 0
	for _, value := range slice {
		wg.Add(1)
		go sumOfSquares(number, response, &wg)
		number <- value
		total += <-response
	}
	defer close(number)
	defer close(response)

	fmt.Println(total)
	wg.Wait()
}

func sumOfSquares(number chan int, response chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-number
	num *= num
	response <- num
}
