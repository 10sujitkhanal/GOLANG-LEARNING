package main

import "fmt"

func say(msg string, ch chan string) {
    ch <- msg // send message to channel
}

func main() {
    ch := make(chan string)

    go say("Hello from Goroutine", ch)

    message := <-ch // receive message from channel
    fmt.Println(message)
}


// A channel is a pipe that connects goroutines so they can send and receive data.

// Channels are used for communication between goroutines.

// You can think of it like one goroutine puts data into the pipe, and another takes data out.