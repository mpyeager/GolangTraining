// Level 10: Exercise 2
// Get the code below to run properly.

package main

import "fmt"

// func main()  {
// 	cs := make(chan<- int)

// 	go func() {
// 		cd <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("-------\n")
// 	fmt.Printf("cs\t%T\n", cs)

// }

func main()  {
	channelSend := make(chan int)

	go func() {
		channelSend <- 42
	}()
	fmt.Println(<-channelSend)

	fmt.Printf("-------\n")
	fmt.Printf("channelSend\t%T\n", channelSend)

}

// func main()  {
// 	cr := make(<-chan int)

// 	go func() {
// 		cr <- 42
// 	}()
// 	fmt.Println(<-cr)

// 	fmt.Printf("-------\n")
// 	fmt.Printf("cr\t%T\n", cr)

// }

func main()  {
	channelReceive := make(chan int)

	go func() {
		channelReceive <- 42
	}()
	fmt.Println(<-channelReceive)

	fmt.Printf("-------\n")
	fmt.Printf("channelReceive\t%T\n", channelReceive)

}
