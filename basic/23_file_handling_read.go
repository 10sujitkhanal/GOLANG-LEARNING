package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("example.txt") // reads the whole file
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(data)) // convert bytes to string
}
