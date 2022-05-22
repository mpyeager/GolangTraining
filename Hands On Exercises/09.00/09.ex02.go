// Level 9: Exercise 2
// Create a type person struct and attach a method speak to type person using a pointer receiver. Create a type human interface. Create func saySomething and have it take in a human as a parameter which calls the speak method.

package main

import "fmt"

type rabbi struct {
	surname string
}

type human interface {
	speak()
}

func (r *rabbi) speak()  {
	fmt.Println("That which is hateful to you, do not do unto your fellow.")
}

func saySomething(h human)  {
	h.speak()
}

func main()  {

	r1 := rabbi{
		surname: `Hillel`
	}

	saySomething(&r1)

	// This won't work as no pointer
	// saySomething(r1)

	// This will also work
	r1.speak()
}
