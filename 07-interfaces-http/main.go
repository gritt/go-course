package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// creates an empty 99999 byte slice to be filled by the Read()
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// turn down for what!
	io.Copy(lw, resp.Body)
}

// logWriter type is now implemeting the Write interface!
func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))

	return len(bs), nil
}
