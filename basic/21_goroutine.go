package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello")
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go sayHello()

    for i := 0; i < 5; i++ {
        fmt.Println("World")
        time.Sleep(100 * time.Millisecond)
    }
}


// A goroutine is a lightweight thread managed by the Go runtime.

// It lets you run functions concurrently (at the same time) without manually creating threads.

// Goroutines are very cheap in memory compared to traditional threads.

// When you call a function using the keyword go, it runs independently in the background.