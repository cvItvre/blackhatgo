package main

import (
	"fmt"
	"net"
)

func main() {

	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		fmt.Printf("Scanning port %d...\n", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered
			fmt.Printf("Port %d is closed or filtered...\n", port)
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open !!!\n", port)
	}

}
