// Level 4: Exercise 5
// To delete from a slice, use append along with slicing. Write a program following these steps; start with slice x := []int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}, use append and slicing to derive values [28, 29, 30, 34, 35, 36, 37] and print.

package main

import "fmt"


func main()  {
	x := []int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Println(x)
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
