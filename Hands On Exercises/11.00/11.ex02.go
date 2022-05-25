// Level 11: Exercise 2
// Beginning with the code below, create a custom error message using fmt.Errorf

package main

import (
	"encoding/json"
	"fmt"
)

type rabbi struct {
	Surname					string
	Forname 				string
	PublishedWorks	[]string
}

func main()  {
	r1 := rabbi{
		Surname: 				"Soloveitchik",
		Forname: 				"Joseph Ber",
		PublishedWorks: []string{"The Lonely Man of Faith", "Halakhic Man", "Halachic Mind"},
	}

	bs, err := toJSON(r1)
	fmt.Println(string(bs))



}
