// Level 6: Exercise 4
// Create a user defined struct with identifer person and fields forname, surname, age. Attach a method to type person with identifier speak, and the method should take the form of person stating their name, age. Create a value of type person, and call the method from value of type person.

package main

import "fmt"

type person struct {
	forname string
	surname string
	age int
}

func (p person) speak()  {
	fmt.Println(`Ich bin`, p.forname, p.surname, `und Ich bin `, p.age, `Jahre alt.`)
}
func main()  {
	p1 := person{
		forname: `Wolfgang Amadeus`,
		surname: `Mozart`,
		age: 35,
	}
	p1.speak()
}
