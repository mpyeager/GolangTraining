// Level 7: Exercise 2
// Create a person struct. Create a func `changeMe` with a *person as a parameter. Change the value stored at *person address. Create a value of type person and print out the value. In func main, call `changeMe` and print out the value.

package main

import "fmt"

// I learn best by creating my own examples, so I've modified the instructions to change the counting of the Omer as `omerState`. Today is 19 May 2022 / 18 Iyar 5782, and Lag B'Omer and the rest of the code should be self explanatory. :-)

type omerDay struct {
	omerDay string
}

func main() {
	omerState := omerDay{
		omerDay: `Lag B'Omer`,
	}
	fmt.Println(omerState)
	updateOmer(&omerState)
	fmt.Println(omerState)
}

func updateOmer(o *omerDay) {
	o.omerDay = `Day 34`
	(*o).omerDay = `Day 35`
}
