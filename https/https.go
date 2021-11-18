package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "helloworld")
    })
    http.ListenAndServeTLS("localhost:8089", "ca.crt", "server.key", nil)
}
