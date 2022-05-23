// Level 10: Exercise 7
// Write a program that; launches 10x goroutines where each goroutine adds 10x numbers to a channel, pulls the numbers off the channel and prints them.

package main

import "fmt"

func main() {
	channel := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				channel <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-channel)
	}

	fmt.Println("About to exit.")
}
