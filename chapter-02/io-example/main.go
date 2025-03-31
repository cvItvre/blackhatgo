package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	// Read some data, from somewhere to anywhere
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	// Write data somewhere
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {

	var reader FooReader
	var writer FooWriter

	// Create a buffer to hold input/output
	input := make([]byte, 4096)

	// Read the input
	b, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data.")
	}
	fmt.Printf("Read %d bytes from stdin\n", b)

	// Write the output
	b, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", b)

}
