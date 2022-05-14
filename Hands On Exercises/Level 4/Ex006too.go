// Level 4: Exercise 6
// Create a slide to store the names of all 25 finalists of Eurovision 2022. Use make and append to do this. Goal: Do not have the array that underlies the slice created more than once. Print all values including index position without using the range clause.

package main

import (
	"fmt"
)

func main() {

	x := make([]string, 0, 25)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	EurovisionFinalists := []string{`Czech Republic`, `Romania`, `Portugal`, `Finland`, `Switzerland`, `France`, `Norway`, `Armenia`, `Italy`, `Spain`, `Netherlands`, `Ukraine`, `Germany`, `Lithuania`, `Azerbaijan`, `Belgium`, `Greece`, `Iceland`, `Moldova`, `Sweden`, `Australia`, `United Kingdom`, `Poland`, `Serbia`, `Estoania`}
	x = append(x, EurovisionFinalists...)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}