// Level 9: Exercise 6
// Create a program that prints out your OS and ARCH.

package main

import (
	"fmt"
	"runtime"
)

func main()  {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
