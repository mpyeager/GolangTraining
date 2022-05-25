// Level 11: Exercise 1
// Beginning with the code below, instead of using the blank identifier ensure the code is checking and handling the error.

package main

import (
	"encoding/json"
	"fmt"
)

type rabbi struct {
	Surname		string
	Forname 	string
	PublishedWorks	[]string
}

func main()  {
	r1 := rabbi{
		Surname: 				"Soloveitchik",
		Forname: 				"Joseph Ber",
		PublishedWorks: []string{"The Lonely Man of Faith", "Halakhic Man", "Halachic Mind"},
	}

	// bs, _ := json.Marshal(r1)
	// fmt.Println(string(bs))

	bs, err := json.Marshal(r1)
	if err != nil {
		log.Fatalln("JSON was not Marshaled, error is: ", err)
	}
	fmt.Println(string(bs))

}
