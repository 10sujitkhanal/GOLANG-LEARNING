package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()

    if _, err := f.WriteString("\nAdding new line!"); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Appended successfully")
    }
}
