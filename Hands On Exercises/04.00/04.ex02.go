// Level 4: Exercise 2
// Using a composite literal; create a slice of type int, assign 10 values. Range over the slice and print the values out. Using format printing, print out the type of the slice.

package main

import "fmt"


func main() {
	x := []int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}
	for i, v := range x {
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n", x)
}
