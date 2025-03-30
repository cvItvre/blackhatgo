package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for range cap(ports) {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)

}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {

		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("Port %d is open!\n", p)
			conn.Close()
		}

		wg.Done()
	}
}
