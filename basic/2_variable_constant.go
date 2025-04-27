package main

import "fmt"

func main() {
	var name string = "Sujit"
	var age int = 30
	city := "Pokhara"

	const PI = 3.14

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("City: ", city)
	fmt.Println("PI: ", PI)
}


// var is for declaring variables with type.

// := is short for declaring and initializing (only inside functions).

// const defines constants.