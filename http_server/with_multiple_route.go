package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome Home!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is About Page!")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
