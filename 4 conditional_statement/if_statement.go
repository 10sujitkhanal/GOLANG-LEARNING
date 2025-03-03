package main

import "fmt"

func main() {
	// variable declaration
	age := 15

	// if statement
	if_statement(age)

	// if else statement
	if_else_statement(age)

	// if else if statement
	if_else_if_statement(age)

}

func if_statement(age int) {
	if age >= 18 {
		fmt.Println("You can vote")
	}
}

func if_else_statement(age int) {
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}
}

func if_else_if_statement(age int) {
	if age >= 18 {
		fmt.Println("You can vote")
	} else if age >= 16 {
		fmt.Println("You can't vote")
	} else {
		fmt.Println("You are too young")
	}
}
