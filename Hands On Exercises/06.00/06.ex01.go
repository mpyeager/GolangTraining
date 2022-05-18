// Level 6: Exercise 1
// Create a function with the identifier foo that returns an int. Create a function with the identifier bar that returns an int and a string. Call both functions and print out their results.

package main

import "fmt"

func main()  {
	bd := foo()
	isr, founding := bar()

	fmt.Println(bd, isr, founding)
}
	func foo() int {
		return 1917
	}

	func bar () (int, string)  {
		return 1948, "Am Yisrael Chai"
}

