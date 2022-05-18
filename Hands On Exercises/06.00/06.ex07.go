// Level 6: Exercise 7
// Assign a function to a variable and call that function

package main

import "fmt"

func main() {

	// As this year marks Her Majesty's Platinum jubilee, let's show our respect. :-)
qe2Reign :=	func() {
		for qe2 := 1952; qe2 < 2022; qe2++ {
			fmt.Println(qe2, `God save the Queen!`)
		}
	}
	qe2Reign()
	fmt.Println(`70 years on the throne!`)
}
