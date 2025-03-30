package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(port int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				fmt.Printf("Port %d is open...\n", port)
				conn.Close()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

}
