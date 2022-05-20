// Level 1: Exercise 4
// Create your own `type``. Have the underlying `type` be an `int`. Create a `variable` of your new `type` with the identifier `x` using the `var` keyword.

package main

import "fmt"

type omer int

var x omer

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Assigning current day of Omer for 09.May.2022
	x = 23
	fmt.Println(x)

}
