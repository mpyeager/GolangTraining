// Level 3: Exercise 9
// Create a program that uses a switch statement with the switch expression specified as a variable of type string with the identifier fav[Something]

package main

import "fmt"

// As we've been counting the Omer, let's stick with Hagim for fav[Something] :-)
func main() {
	favHag := "Pesach"
	switch favHag {
	case "Pesach":
		fmt.Println("No one likes matzah _that_ much!")
	case "Shavuot":
		fmt.Println("I mean, yeah ... cheescake is life!")
	case "Sukkot":
		fmt.Println("Don't forget to take your jacket to the Sukkah!")
	case "Rosh Hashana":
		fmt.Println("You can never have too many fish balls. Wait ... fish have balls?!")
	case "Yom Kippur":
		fmt.Println("Nil by mouth.")
	}
}
