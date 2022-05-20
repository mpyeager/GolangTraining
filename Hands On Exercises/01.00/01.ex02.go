// Level 1: Exercise 3
// Using the code from the previous exercise; at the package level scope, assign the values to three `variable`s. In `func main` use `fmt.Sprintf` to print all of the values to a single `string`. Assign the returned value of `type string` using the short declaration operator to a `variable` with the identifier `s`.


package main

import "fmt"

var x int = 42
var y string = "Bond. James Bond."
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

}
