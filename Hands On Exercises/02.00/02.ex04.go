// Level 2: Exercise 4
// Write a program that; assigns an int to a variable, prints that int in [decimal, binary, and hex], shifts the bits of that int over 1 position to the left and assigns that to a variable, prints that variable in [decimal, binary, and hex].

package main

import "fmt"

func main() {
	Omer := 24
	fmt.Printf("%d\t%b\t%#x\n", Omer, Omer, Omer)

	//Sift left by 1.
	OmerShift := Omer << 1
	fmt.Printf("%d\t%b\t%#x\n", OmerShift, OmerShift, OmerShift)
}
