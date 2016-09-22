package main

import (
  "log"
  "net/http"
  "runtime"
)

func main() {
  //runtime.GOMAXPROCS(runtime.NumCPU() - 1)
  runtime.GOMAXPROCS(6)

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Last-Modified", "Thu, 18 Jun 2015 10:24:27 GMT")
    w.Header().Add("Accept-Ranges", "bytes")
    w.Header().Add("E-Tag", "55829c5b-17")
    w.Header().Add("Server", "golang-http-server")
    w.Write([]byte("Hello world!\n"))
  })

  log.Printf("%d", runtime.NumCPU())
  log.Printf("Go http Server listen on :12345")
  log.Fatal(http.ListenAndServe(":12345", nil))
}

