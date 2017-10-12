package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// second index is the file path
	fp := os.Args[1]

	// first param implmenets Writer, second param implement Reader
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}

	// reads input, copy (write) to output
	io.Copy(os.Stdout, f)
}
