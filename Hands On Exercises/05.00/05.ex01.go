// Level 5: Exercise 1
// Create your own type person with will have an underlying type of struct such that it can store the following data; first name, last name, favourite 80s cartoon characters. Create two values of type person and print out the values, ranging over the elements in the slice.

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

	fmt.Println(p1.surname, "," ,p1.first)
	for i, v := range p1.favCartChars {
		fmt.Println(i, v)
	}

	fmt.Println(p2.surname, ",", p2.first)
	for i, v := range p2.favCartChars {
		fmt.Println(i, v)
	}

	fmt.Println(p3.surname, ",", p3.first)
	for i, v := range p3.favCartChars {
		fmt.Println(i, v)
	}
}
// Level 5: Exercise 1
// Create your own type person with will have an underlying type of struct such that it can store the following data; first name, last name, favourite 80s cartoon characters. Create two values of type person and print out the values, ranging over the elements in the slice.

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

	fmt.Println(p1.surname,`,`,p1.first)
	fmt.Println(`Favourite 80s Cartoon Characters:`)
	for i, v := range p1.favCartChars {
		fmt.Println(i, v)
	}

	fmt.Println(p2.surname,`,`,p2.first)
	fmt.Println(`Favourite 80s Cartoon Characters:`)
	for i, v := range p2.favCartChars {
		fmt.Println(i, v)
	}

	fmt.Println(p3.surname,`,`,p3.first)
	fmt.Println(`Favourite 80s Cartoon Characters:`)
	for i, v := range p3.favCartChars {
		fmt.Println(i, v)
	}
}
