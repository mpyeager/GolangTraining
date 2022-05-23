// Level 10: Exercise 5
// Show the comma ok idiom starting with the code below.

package main

import "fmt"

// func main() {
// 	c := make(chan int)

// 	v, ok := <-c
// 	fmt.Println(v, ok)

// 	close(c)

// 	v, ok = <-c
// 	fmt.Println(v, ok)
// }

func main() {
	c := make(chan int)

	go func() {
		c <-42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
