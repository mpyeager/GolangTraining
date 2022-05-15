// Level 2: Exercise 6
// Using iota, create 4 constants for the last 4 years and print the constant values.

package main

import "fmt" "fmt"

// Since we've been counting the Omer, let's change the exercise a little bit and count that using iota instead. :-)
const (
	// We use _ to increment the counting from zero as the Omer doesn't start on a Day Zero. :-)
	_            = iota
	OmerDayOne   = 0 + iota
	OmerDayTwo   = 0 + iota
	OmerDayThree = 0 + iota
	OmerDayFour  = 0 + iota
)

func main() {
	fmt.Println(OmerDayOne)
	fmt.Println(OmerDayTwo)
	fmt.Println(OmerDayThree)
	fmt.Println(OmerDayFour)
}
