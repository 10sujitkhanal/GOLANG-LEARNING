package main

import "fmt"

func addSubtract(x int, y int) (int, int) {
    return x + y, x - y
}

func main() {
    sum, diff := addSubtract(10, 5)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
