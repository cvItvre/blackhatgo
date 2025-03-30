package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"tcp-scanner-user/portformat"
	"time"
)

func main() {

	var ports []int
	var args string

	fmt.Print("Which ports do you want to scan: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	args = scanner.Text()

	ports, err := portformat.Parse(args)
	if err != nil {
		log.Fatalln(err)
	}

	for _, port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)

		conn, err := net.DialTimeout("tcp", address, time.Second*5)
		if err != nil {
			continue
		}
		fmt.Printf("Port %d open!\n", port)
		conn.Close()
	}

}
