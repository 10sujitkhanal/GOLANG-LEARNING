package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you requested: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at port 8080...")
    http.ListenAndServe(":8080", nil)
}
