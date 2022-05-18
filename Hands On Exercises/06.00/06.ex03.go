// Level 5: Exercise 3
// Use the defer keyword to show that a deferred function runs after the function containing it exits.

package main

import "fmt"

func main()  {
	defer motziShabbat()
	fmt.Println(`Shabbat Shalom!`)
}

func motziShabbat() {
	fmt.Println(`Shavua Tov!`)
}
