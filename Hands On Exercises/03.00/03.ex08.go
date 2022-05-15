// Level 3: Exercise 8
// Create a program that uses a switch statement with no switch expression specified

package main

import "fmt"


func main() {
	switch {
	case false:
		fmt.Println("We are not counting the Omer.")
	case true:
		fmt.Println("We are counting the Omer.")
	}

}
