// Level 5: Exercise 2
// Create a function with the identifier foo that; takes in a variadic parameter of type int, pass in a value of type []int to the function, returns the sum of all values of type int passed in. Create a function with the identifier bar that; takes in a parameter of type []int, returns the sum of all values of type in passed in.

package main

import "fmt"

func main()  {
	ii := []int{1, 2, 3, 4, 67, 93, 428, 911, 1948}
	n := foo(ii...)
	fmt.Println(n)

	iii := []int{1, 2, 3, 4, 67, 93, 428, 911, 1948}
	ntoo := bar(iii)
	fmt.Println(ntoo)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
