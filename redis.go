package main

import (
    //"io"
    "runtime"
    "fmt"
    "net/http"
    "log"
    "gopkg.in/redis.v4"
)

var client = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

func redisTest(w http.ResponseWriter, req *http.Request) {
    val, err := client.Get("key").Result()
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println("key", val)
    w.Write([]byte (val))
}

func main() {
    runtime.GOMAXPROCS(6)

    http.HandleFunc("/redis", redisTest)
    log.Fatal(http.ListenAndServe(":12345", nil))
}
