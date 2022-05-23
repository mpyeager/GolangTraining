// Level 10: Exercise 3
// Get the code below to run properly.

package main

import "fmt"

// func main()  {
// 	c := gen()
// 	receive(c)

// 	fmt.Println("About to exit.")
// }

// func gen() <-chan int {
// 	c := make(chan int)

// 		for i := 0; i < 100; i++ {
// 			c <-i
// 		}
// 	}()
// 	return c
// }

func main()  {
	c := gen()
	receive(c)

	fmt.Println("About to exit.")
}

func receive(c <-chan int)  {
	for v := range c {
		fmt.Println(v)
	}
}
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <-i
		}
		close(c)
	}()

	return c
}
