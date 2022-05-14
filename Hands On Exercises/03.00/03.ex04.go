// Level 3: Exercise 4
// Create a for loop using syntax for {}, print the years you have been alive

package main

import (
	"fmt"
)

func main() {
	bd := 1974
	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}