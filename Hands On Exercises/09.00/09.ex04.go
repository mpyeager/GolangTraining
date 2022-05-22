// Level 9: Exercise 4
// Fix the race condition you created in the previous exercise by using a mutex.

package main

import (
	"fmt"
	"sync"
)
func main()  {

	var waitgroups sync.WaitGroup

	incrementer := 0
	gs := 100
	waitgroups.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			waitgroups.Done()
		}()
	}
waitgroups.Wait()
fmt.Println("Ending Value:", incrementer)
}
