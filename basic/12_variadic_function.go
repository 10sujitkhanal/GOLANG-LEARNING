package main

import "fmt"

func sum(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Sum:", total)
}

func main() {
	sum(1, 2, 3)         // Sum: 6
	sum(10, 20, 30, 40)  // Sum: 100
	sum()                // Sum: 0
}
