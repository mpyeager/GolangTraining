// Level 4: Exercise 4
// Write a program following these steps; start with the slice x :=[]int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}, append 38 to the slice and print, in one statement append 39, 40, 41 and print, append y := []int{42, 43, 44, 45, 46} and print.

package main

import "fmt"


func main()  {
	x :=[]int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}
	x = append(x, 38)
	fmt.Println(x)
	x = append(x, 39, 40,41)
	fmt.Println(x)

	y := []int{42, 43, 44, 45, 46}
	x = append(x, y...)
	fmt.Println(x)
}
