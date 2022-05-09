// We can ask Go to tell us exactly what operating system and architecture
// type we're running. :-)

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
