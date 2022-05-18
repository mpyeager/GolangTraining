// Level 6: Exercise 6
// Build and use an anonymous function.

// As this year marks Her Majesty's Platinum jubilee, let's show our respect. :-) We'll use this exercise to show the years Queen Elizabeth II has reigned.

package main

import "fmt"

func main() {

	func() {
		for qe2 := 1952; qe2 < 2022; qe2++ {
			fmt.Println(qe2, `God save the Queen!`)
		}
	}()
}
