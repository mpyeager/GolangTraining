// Level 3: Exercise 2
// Print every rune code point of the uppercase ASCII alphabet 3x

package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		
	} 
}