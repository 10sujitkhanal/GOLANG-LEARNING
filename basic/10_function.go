package main

import "fmt"


// No return function
func greet(name string) {
	fmt.Println("Hello", name)
}

// This is a function named greet.

// It takes one input: name (a string).

// It does not return anything — it only prints Hello <name>.

// Example: If you pass "Sujit", it will print Hello Sujit





// Return function
func greet2(name string) string{
	return name
}

// This is another function named greet2.

// It also takes one input: name (a string).

// BUT this one returns a string (notice the string after the parentheses).

// It simply returns the name back to the caller.




func main() {
	greet("Sujit") //call no return function
	message := greet2("Sujit") //call return function
	fmt.Println(message)
}