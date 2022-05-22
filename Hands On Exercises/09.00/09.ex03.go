// Level 9: Exercise 3
// Using goroutines, create an incrementer program. Have a variable to hold the incrementer value, and launch several goroutines with each goroutine; reading the incrementer value, store in a new variable, yield the processor with runtime.Gosched(), increment the new variable, write the value in the new variable back to the incrementer variable. Use waitgroups to ensure all goroutines finish. Prove that you've just created a race condition using the -race flag.

package main

import (
	"fmt"
	"sync"
	"runtime"
)
func main()  {

	var waitgroups sync.WaitGroup

	incrementer := 0
	gs := 100
	waitgroups.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			waitgroups.Done()
		}()
	}
waitgroups.Wait()
fmt.Println("Ending Value:", incrementer)
}
