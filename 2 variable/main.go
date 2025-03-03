package main

import "fmt"

func main() {
	// Declaration with var
	var age int = 30
	var name string = "Sujit"
	var isActive bool = true

	// Short declaration with type inference
	score := 90.5

	// Array
	var arr [3]int = [3]int{1, 2, 3}

	// Slice
	slice := []int{4, 5, 6, 7}

	// Printing the variables
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Is Active:", isActive)
	fmt.Println("Score:", score)
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}
