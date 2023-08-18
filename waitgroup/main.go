package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	endPoint := []string{
		"https://google.com",
		"https://fb.com",
		"https://youtube.com",
	}
	for _, url := range endPoint {
		go getEndPoint(url)
		wg.Add(1)
	}

	wg.Wait()
}

func getEndPoint(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("Failed to get the endpoint", err)
	} else {
		fmt.Printf("%d is the status code and the result is %s", res.StatusCode, endPoint)
		fmt.Println("")
	}
}
