// Level 6: Exercise 10
// Closure is defined as enclosing the scope of a variable in a code block. Create a function and enclose the scope of a variable.

// As this year marks Her Majesty's Platinum jubilee, let's show our respect. :-) We'll use this exercise to calculate the coronation year of Queen Elizabeth II.

package main

import "fmt"

func main() {

	qe2Coronation := yew()
	fmt.Println(qe2Coronation())
}

func yew() func() int {
	gviDeath := 1951
	return func() int {
		gviDeath++
		return gviDeath
	}
}
