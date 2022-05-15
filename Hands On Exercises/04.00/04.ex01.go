// Level 4: Exercise 1
// Using a composite literal; create and array which holds 5 valuse of type int, assigne values to each index position. Range over the array and print the values out. Using format printing, print out the type of the array.

package main

import "fmt"


func main() {
	x := [5]int{27, 28, 29, 30, 31}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
