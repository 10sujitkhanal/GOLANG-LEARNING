package	main

import "fmt"

func main() {
	for i:=1; i<=10; i++ {
		if i == 5 {
			continue // Skip 5
		}
		if i == 8 {
			break // Stop at 8
		}
		fmt.Println(i)
	}
}