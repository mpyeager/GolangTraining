// Exercise 4
// Create your own type. Have the underlying type be an int.
// Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the VAR keyword

package main

import (
	"fmt"
)

type omer int

var x omer

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// Assigning current day of Omer for 09.May.2022
	x = 23
	fmt.Println(x)

}
