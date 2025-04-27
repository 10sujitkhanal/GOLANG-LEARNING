package main 
import "fmt"

func main(){
	var name string
	var age int
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("Hello ", name, ", you are ", age, " years old!")
}