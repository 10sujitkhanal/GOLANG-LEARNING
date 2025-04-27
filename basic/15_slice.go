package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}

    fruits = append(fruits, "Mango") // Add element
    fmt.Println(fruits) // [Apple Banana Cherry Mango]

    fmt.Println(fruits[1:3]) // [Banana Cherry] (slice)
}
