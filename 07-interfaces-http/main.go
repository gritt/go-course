package main

import (
  "net/http"
  "os"
  "fmt"
)

func main () {

    resp, err := http.Get("http://google.com")

    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
    }

    b := []byte

    fmt.Println(resp.Body.Read(b))
}
