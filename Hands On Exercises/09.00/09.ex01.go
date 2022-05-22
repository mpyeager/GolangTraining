// Level 9: Exercise 1
// In addition to the main goroutine, launch 2x additional goroutines with each net additional goroutine printing something out. Use waitgroups to ensure each goroutine finishes before your program exits.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Beginning Execution CPUs:\t", runtime.NumCPU())
	fmt.Println("Beginning Execution GS:\t\t", runtime.NumGoroutine())
	fmt.Println("----------")

	var waitgroups sync.WaitGroup
	waitgroups.Add(2)

	go func() {
		fmt.Println("Hello from Thing One!")
		waitgroups.Done()
	}()

	go func() {
		fmt.Println("Hello from Thing Two!")
		waitgroups.Done()
	}()

	fmt.Println("Mid Execution CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Mid Execution GS:\t\t", runtime.NumGoroutine())
	fmt.Println("----------")

	waitgroups.Wait()

	fmt.Println(" << Program exiting. >> ")
	fmt.Println("----------")
	fmt.Println("End Execution CPUs:\t\t", runtime.NumCPU())
	fmt.Println("End Execution GS:\t\t", runtime.NumGoroutine())

}
