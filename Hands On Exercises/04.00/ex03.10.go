// Level 4: Exercise 10
// Using the code from the previous exercise, delete a record from the map. Print the map using range loop.

package main

import (
	"fmt"
)

func main()  {
	bp := map[string][]string{
		`Turing, Alan`: []string{`Creator of Enigma "Prof's Book"`, `Developed statistical procedure 'Banburismus' to increase Bombe machine efficiency`, `Created 'Applications of Probability to Cryptography and 'Paper on Statistics of Repetitions'`},
		`Welchman, Gordon`: []string{`Member of the "Wicked Uncles"`, `Developed metadata traffic analysis process`, `Head of Hut 6, Enigma cipher break`},
		`Batey, Mavis`: []string{`Hut 6 codebreaker`, `Battle of Matapan decryption`, `Abwehr Enigma message decoding`},
		`Fawcett, Jane`:	[]string{`Hut 6 codebreaker`, `Decoded message detailing Bismarck position`, `Secretary, Victorian Society` },
		`Valentine, Jean`:	[]string{`Bombe operator`, `WRNS`, `Bombe reconstruction, post war`},
	}

	fmt.Println("ALL RECORDS")
	for key, value := range bp {
		fmt.Println("Bletchley Personnel Record: ", key)
		for index, value := range value {
			fmt.Println("\t", index, value)
		}
	}

	fmt.Print("\n")
	fmt.Println(" ... Deleting classified record(s) ... ")
	fmt.Print("\n")
	delete(bp, `Turing, Alan`)

	fmt.Println("DECLASSIFIED")
	for key, value := range bp {
		fmt.Println("Bletchley Personnel Record: ", key)
		for index, value := range value {
			fmt.Println("\t", index, value)
		}
	}

	fmt.Print("\n")
	fmt.Println(" ... Adding declassified record(s) ... ")
	fmt.Print("\n")
	bp[`Watters, Jean Briggs`] = []string{`Ultra programme, Enigma`, `Cryptanalyst`, `WRNS`}

	fmt.Println("DECLASSIFIED COMPLETE")
	for key, value := range bp {
		fmt.Println("Bletchley Personnel Record: ", key)
		for index, value := range value {
			fmt.Println("\t", index, value)
		}
	}
}
