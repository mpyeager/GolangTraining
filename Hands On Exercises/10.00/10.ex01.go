// Level 10: Exercise 1
// Make the code below work. Use a func literal [anonymous self-executing func] as well as a buffered channel.

package main

import "fmt"

// func main()  {
// 	c := make(chan int)

// 	c <- 42

// 	fmt.Println(<-c)
// }

// Anonymous func
func main()  {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Buffered channel
func main()  {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
