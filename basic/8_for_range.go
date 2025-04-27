package main

import "fmt"

func main() {

	fruits := []string{"apple", "banana", "cherry"} //array of strings

	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}