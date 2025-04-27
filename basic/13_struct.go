package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Method (Value Receiver)
func (p Person) greet() {
    fmt.Println("Hello,", p.name)
}

func main() {
    p := Person{name: "Alice", age: 30}
    p.greet()
}
