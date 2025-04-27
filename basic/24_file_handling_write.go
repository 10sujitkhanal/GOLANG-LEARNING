package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File written successfully")
}
