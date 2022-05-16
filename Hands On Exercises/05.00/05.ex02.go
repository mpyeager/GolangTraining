// Level 5: Exercise 2
// Using the code from the previous exercise, store the values of type person in a map with key of surname. Access each value in the map printing the values.

package main

import "fmt"

type person struct {
	first        string
	surname      string
	favCartChars []string
}

func main() {
	p1 := person{
		first:   `Alan`,
		surname: `Turing`,
		favCartChars: []string{
			`Teenage Mutant Ninja Turtles`,
			`Thundercats`,
			`Inspector Gadget`,
		},
	}
	p2 := person{
		first:   `Gordon`,
		surname: `Welchman`,
		favCartChars: []string{
			`He-Man`,
			`Bugs Bunny`,
			`Smurfs`,
		},
	}

	p3 := person{
		first:   `David`,
		surname: `Spiegelhalter`,
		favCartChars: []string{
			`Transformers`,
			`Garfield`,
			`The Simpsons`,
		},
	}

	m := map[string]person{
		p1.surname: p1,
		p2.surname: p2,
		p3.surname: p3,
	}

for _, value := range m {
	fmt.Println(value.surname, value.first)
	for index, value := range value.favCartChars {
		fmt.Println(index, value)
	}
	fmt.Println(`----------`)
}

}
