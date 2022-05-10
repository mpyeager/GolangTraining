// Level 2: Exercise 2
// Using the following operators, write expressions and assign their values to variables; == <= >= != < >

package main

import (
	"fmt"
)

// The exercise doesn't ask for it, but let's use package level variables instead of assigning and initialising each time and, as we're currently couting the Omer, let's use that. :-)

var Omer int = 24
var EndOmer int = 50

func main() {

	a := (Omer == EndOmer)
	b := (Omer <= EndOmer)
	c := (Omer >= EndOmer)
	d := (Omer != EndOmer)
	e := (Omer < EndOmer)
	f := (Omer > EndOmer)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
