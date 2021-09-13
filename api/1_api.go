package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-chi/chi"
    "log"
    "net/http"
)

func main() {
    fmt.Println("start server")
    router := chi.NewRouter()
    router.Get("/api/getExample", getHandler)
    router.Post("/api/postExample", PostHandler)

    log.Fatal(http.ListenAndServe(":8091", router))
    fmt.Println("server is listening on port 8091")
}

func getHandler(w http.ResponseWriter, r *http.Request)  {
    json.NewEncoder(w).Encode("you got me")
}

func PostHandler(w http.ResponseWriter, r *http.Request)  {
    json.NewEncoder(w).Encode("you got me")
}
