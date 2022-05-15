// Level 3: Exercise 7
// Building on the previous exercise, create a program that uses else if and else

package main

import "fmt"
// As we've been counting the Omer, let's stick with that. :-)
func main() {
	Omer := 66

	if Omer < 50 {
		fmt.Printf("Today is day %v of the Omer.\nWe have not reached the end of the counting of the Omer.", Omer)
	} else if Omer == 50 {
		fmt.Printf("Today is day %v of the Omer.\nWe have reached the end of the counting of the Omer.", Omer)
	} else {
		fmt.Println("We are not counting the Omer.")
	}
}
