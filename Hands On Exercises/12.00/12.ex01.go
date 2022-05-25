// Level 12: Exercise 1
// Create a package dog where the package should have an expected func years whihc converts human years to dog years. Document your code with comments, then run program to ensure it is functional and then run local server with godoc to view documentation.

package main

import "fmt"

type hund struct {
	name 	string
	age 	int
}

func main() {
	dog := hund{
		name: "Rafi",
		age: dog.Years(10),
	}
	fmt.Println(dog)
}
