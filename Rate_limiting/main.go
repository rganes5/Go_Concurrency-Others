package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

/*
The client code sends a total of 10 requests in batches of 3 (controlled by max).
It creates multiple goroutines to send requests concurrently. The client connects to the server on port 8080 using net.Dial and reads the response.
If the response is not "success," it logs an error. The client simulates multiple clients sending requests to the server.
*/

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8081")
				if err != nil {
					log.Fatalf("could not dial from main: %v", err)
				}
				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("could not read from conn: %v", err)
				}
				if string(bs) != "success" {
					log.Fatalf("request error, request: %d", i+1+j)
				}

				fmt.Printf("request %d: success\n", i+1+j)
			}(j)
		}
		wg.Wait()
	}
}
