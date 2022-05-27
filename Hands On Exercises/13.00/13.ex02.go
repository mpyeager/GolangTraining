// Level 13: Exercise 2
// Beginning with the code below, get the code ready to BET; Benchmarks, Examples, Tests. However, do not write an example for the func which returns a map, rather, write a test as an extra challenge. Add documentation to the code and run in the following order; test(s), benchmark(s), coverage, coverage (web browser), examples (documentation in web browser).

package main

import (
	"fmt"
	"word"
	"poem"
)

func main() {
	fmt.Println(word.Count(poem.SongOfMyself))

	for _, v := range word.UseCount(poem.SongOfMyself) {
		fmt.Println(v, k)
	}
}
