package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Fprintf 格式化输出到w
        fmt.Fprintf(w, "hello")
    })
    http.ListenAndServe("localhost:8089", nil)
}
