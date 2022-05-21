// Level 8: Exercise 4
// Sort the int[] and string[].

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{`we`, `don't`, `make`, `mistakes`, `just`, `happy`, `little`, `accidents`, `there's`, `nothing`, `wrong`, `with`, `having`, `a`, `tree`, `as`, `a`, `friend`}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

