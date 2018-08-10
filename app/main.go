package main

import (
    "net/http"
    "log"
    "fmt"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/html; charset=utf8")
    w.Write([]byte("こんにちは"))
}

func main() {
    fmt.Println("テストです。ほげほげふーばー")
    http.HandleFunc("/", rootHandler)
    log.Fatal(http.ListenAndServe(":80", nil))
}