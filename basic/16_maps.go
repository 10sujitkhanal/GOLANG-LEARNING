package main

import "fmt"

func main() {
    student := map[string]int{
        "Math":    90,
        "Science": 85,
    }

    student["English"] = 88 // add new key-value
    fmt.Println(student)

    fmt.Println(student["Math"]) // 90

    // Deleting a key
    delete(student, "Science")
    fmt.Println(student)
}
