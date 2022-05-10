// Level 1: Exercise 5
// Building on the code from Ex004; at the package level scope, using the VAR keyword, create a VARIABLE with the identifier "y". The VARIABLE should be of the UNDERLYING TYPE of your custom TYPE "x".

package main

import (
	"fmt"
)

type omer int

var x omer
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 23
	fmt.Println(x)

	// Now use CONVERSION to convert the TYPE of the VALUE stored in "x" to the underlying type
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
