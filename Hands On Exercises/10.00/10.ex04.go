// Level 10: Exercise 4
// Starting with the code below, pull the values off the channel using the select statement.

package main

import "fmt"

// func main() {
// 	q := make(chan int)
// 	c := gen(q)

// 	receive(c, q)

// 	fmt.Println("About to exit.")
// }

// func gen(q <-chan int) <-chan int {
// 	c := make(chan int)

// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}

// 	return c
// }

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("About to exit.")
}

func receive(c, q <-chan int)  {
	for {
		select{
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <-1
		close(c)
	}()

	return c
}
