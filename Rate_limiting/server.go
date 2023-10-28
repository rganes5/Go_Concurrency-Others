package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

/*
The server code listens on port 8080 using net.Listen and handles incoming connections.
It limits the number of concurrent connections to 3. When a client connects, the server sends a "success" response after simulating heavy work using time.Sleep.
*/

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("could not create listener: %v", err)
	}

	var connections int32
	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		connections++

		go func(con net.Conn) {
			defer func() {
				_ = con.Close()
				atomic.AddInt32(&connections, -1)
			}()
			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			// simulate heavy work
			time.Sleep(time.Second)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Fatalf("could not write to connection: %v", err)
			}
		}(conn)
	}
}
