// Level 10: Exercise 6
// Write a program that; puts 100x numbers to a channel, pulls the numbers off the channel and prints them.

package main

import "fmt"

func main()  {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
	fmt.Println("About to exit.")
}
