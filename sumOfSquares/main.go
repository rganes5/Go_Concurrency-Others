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

/*

I see the confusion. It seems like you're thinking that the function call `sumOfSquares(number, response, &wg)`
happens before sending the value to the `number` channel, but in reality,
the goroutine starts executing immediately after it's created. Here's the order of execution:

1. A goroutine is created with `go sumOfSquares(number, response, &wg)`.

2. Inside this goroutine, it starts executing the `sumOfSquares` function.

3. As part of the `sumOfSquares` function, it waits to receive a value from the `number` channel with `num := <-number`.

4. In the main goroutine, you send the value `value` into the `number` channel with `number <- value`.

5. The goroutine created in step 1 receives the value from the `number` channel, calculates the square, and sends the squared result to the `response` channel.

So, the function call and the value being sent to the channel do not happen in a sequential manner in the main goroutine.
 Instead, they happen concurrently. This is the power of goroutines in Go – they can execute independently and concurrently, allowing you to perform tasks in parallel.

*/
