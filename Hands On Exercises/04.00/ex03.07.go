// Level 4: Exercise 7
// Create a slice of a slice of string ([][]string). Store data in the multi-dimensional slice. Range over the records and then range of the data in each record.

package main

import (
	"fmt"
)

func main()  {
	xsShammai := []string{"Beit Shammai", "Torah Study", "Only worthy students should be admitted to study."}
	xsHillel := []string{"Beit Hillel", "Torah Study", "May be taught to anyone that they will repent and become worthy."}
	fmt.Println(xsShammai)
	fmt.Println(xsHillel)

	xxs := [][]string{xsShammai, xsHillel}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t Index Position: %v \t Value: %v \n", j, val)
		}
	}
}
