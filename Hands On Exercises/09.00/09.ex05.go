// Level 9: Exercise 5
// Fix the race condition you created in ex03 by using package atomic.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var waitgroups sync.WaitGroup
	var incrementer int64

	gs := 100
	waitgroups.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			waitgroups.Done()
		}()
	}
	waitgroups.Wait()
	fmt.Println("Ending Value:", incrementer)
}
