// Level 1: Exercise 5
// Building on the code from the previous exercise; at the package level scope, using the `var` keyword create a `variable` with the identifier `y`. The `variable` should be of the underlying `type` of your custom `type` `x`. Now use `conversion` to convert the `type` of the value stored in `x` to the underlying `type`.


package main

import "fmt"

type omer int

var x omer
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 23
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
