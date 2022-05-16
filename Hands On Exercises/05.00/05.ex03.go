// Level 5: Exercise 3
// Create a new type vehicle with underlying type struct with fields for doors and colour. Next, create 2x new types with underlying type struct; evSaloon and SUV. Make the SUV a fourWD using bool, and make the evSaloon luxury using bool. Using vehicle, evSaloon, and SUV structs, create a value of type SUV and type evSaloon and assign values to the fields using a composite literal. Print each of these values and a single field from each value.

package main

import "fmt"

type vehicle struct {
	doors int
	colour string
}

type SUV struct {
	vehicle
	fourWD bool
}

type evSaloon struct {
	vehicle
	luxury bool
}
func main()  {
	s := SUV{
		vehicle: vehicle{
			doors: 5,
			colour: `black`,
		},
		fourWD: true,
	}

	ev := evSaloon{
		vehicle: vehicle{
			doors: 4,
			colour: `british racing green`,
		},
		luxury: true,
	}

	fmt.Println(`Vehicle Specs: `, s)
	fmt.Println(`Vehicle Specs: `, ev)
	fmt.Println(`Vehicle Colour: `, s.colour)
	fmt.Println(`Vehicle Colour: `, ev.colour)
}
