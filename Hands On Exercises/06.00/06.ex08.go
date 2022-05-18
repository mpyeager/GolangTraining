// Level 6: Exercise 8
// Create a function which returns a function. Assign the returned function to a variable and call the returned function.

package main

import "fmt"

func main() {

	qe2Reign := qe2()
	fmt.Println(qe2Reign())

}

func qe2() func() int {
	return func() int {
		return 70
	}
}
