// Level 3: Exercise 5
// Print out the modulus for each number between 10 and 100 when divided by 4

package main

import "fmt"

func main() {
	 for i := 10; i <= 100; i++ {
		 fmt.Printf("When we divide %v by 4, the modulus is %v\n", i, i%4)
	 }
}
