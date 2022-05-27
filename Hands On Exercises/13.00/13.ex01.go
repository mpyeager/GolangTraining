// Level 13: Exercise 1
// Beginning with the code below, prepare the code for BET; Benchmarks, Examples, Tests. Run in the following order; test(s), benchmark(s), coverage, coverage (web browser), examples (documentation in web browser).

package main

import (
	"fmt"
	"github.com/mpyeager/GolangTraining/Hands%20On%20Exercises/12.00/kelev"
)

type hund struct {
	name 	string
	age 	int
}

func main() {
	dog := hund{
		name: "Rafi",
		age: kelev.Years(10),
	}
	fmt.Println(dog)
	fmt.Println(kelev.YearsTwo())
}
