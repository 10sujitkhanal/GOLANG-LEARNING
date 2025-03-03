package main

import "fmt"

func main() {
	number := 3
	switch {
	case number < 5:
		fmt.Println("Number is less than 5")
	case number == 5:
		fmt.Println("Number is equal to 5")
	default:
		fmt.Println("Number is greater than 5")
	}
}
