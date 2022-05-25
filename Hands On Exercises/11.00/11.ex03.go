// Level 11: Exercise 3
// Create a struct customErr which implements the builtin.error interface.

package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("The error is: %v", ce.info)
}

func main()  {
	nt := customErr{
		info: "This isn't Earl Grey. Need moar tea.",
	}

	foo(nt)
}

func foo(e error)  {
	fmt.Println("Foo has runneth ...", e)
}
