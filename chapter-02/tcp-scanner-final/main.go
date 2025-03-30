package main

import (
	"fmt"
	"net"
	"slices"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		results <- p
		conn.Close()
	}
}

func main() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for range cap(ports) {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for range 1024 {
		p := <-results
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	close(ports)
	close(results)

	slices.Sort(openPorts)
	for _, port := range openPorts {
		fmt.Println(port, "is open!")
	}

}
