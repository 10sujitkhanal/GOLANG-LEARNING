package main

import "fmt"

func main() {
    var numbers [3]int // array of 3 integers
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30

    fmt.Println(numbers)      // [10 20 30]
    fmt.Println(numbers[1])   // 20
    fmt.Println(len(numbers)) // 3
}
