package main

import "fmt"

func main() {
	s := struct {
		surname    string
		forname    string
		authCodes  map[string]int
		authLevels []string
	}{
		surname: `Turing`,
		forname: `Alan`,
		authCodes: map[string]int{
			`Ultra`:  7192,
			`Enigma`: 4242,
			`Bombe`:  6334,
			`Hut 6`:  9921,
		},
		authLevels: []string{
			`Classified`,
			`Secret`,
			`Top Secret`,
		},
	}

	fmt.Println(s.surname, `,`, s.forname)
	fmt.Println(``)

	fmt.Println(`Authorisation Codes`)
	fmt.Println(`----------`)
	for k, v := range s.authCodes {
		fmt.Println(`Project:`, k, `Auth Code:`, v)
	}

	fmt.Println(``)
	fmt.Println(`Approved Authorisation Levels`)
	fmt.Println(`----------`)
	for index, value := range s.authLevels {
		fmt.Println(index, value)
	}
}


