// Level 7: Exercise 1
// Create a value and assign it to a variable. Print the address of the value.

package main

import "fmt"

func main()  {
	monarch := `Queen Elizabeth II`
	fmt.Println(monarch)
	fmt.Println(&monarch)
}
