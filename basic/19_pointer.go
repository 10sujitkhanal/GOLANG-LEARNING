package main

import "fmt"

func main() {
    var a int = 10
    var p *int = &a

    fmt.Println("Value of a:", a)    // 10
    fmt.Println("Address of a:", p)  // address
    fmt.Println("Value at p:", *p)   // 10 (dereferencing)
}
