package main

import (
  "io"
  "net/http"
  "log"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
  ok := make(chan bool)
  go func() {
    io.WriteString(w, ("hello, world!\n"))
    ok <-true
  }()
  <- ok
}

func main() {
    http.HandleFunc("/hello", helloServer)
    log.Fatal(http.ListenAndServe(":12345", nil))
}
