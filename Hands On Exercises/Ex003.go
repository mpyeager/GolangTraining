//Exercise 3
//Using the code from the previous exercise;
//At the package level scope, assign the values to three variables.

package main

import (
	"fmt"
)

var x int = 42
var y string = "Bond. James Bond."
var z bool = true

func main() {
	//In func main use fmt.Sprintf to print all of the values to a single string.
	//ASSIGN the returned value of TYPE string using the short declaration operator to a
	//VARIABLE with the IDENTIFIER "s"
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
