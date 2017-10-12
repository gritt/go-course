package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fileWriter struct{}

func main() {

	// second index is the file path
	fp := os.Args[1]

	// implements the Writer interface, so it can write the output
	lw := fileWriter{}

	// first param implmenets Writer, second param implement Reader
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}

	// reads input, copy (write) to output
	io.Copy(lw, f)
}

// fileWriter is now implemeting the Writer interface
func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
