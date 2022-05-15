// Level 3: Exercise 6
// Create a program that shows the if statement in action

package main

import "fmt"

// As we've been counting the Omer, let's stick with that. :-)
func main() {
	Omer := 27

	if Omer < 50 {
		fmt.Printf("Today is day %v of the Omer.\nWe have not reached the end of the counting of the Omer.", Omer)
	}
}
