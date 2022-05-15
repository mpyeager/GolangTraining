//Level 4: Exercise 8
// Create a map with a key of type string which is the full name of a Bletchley Park researcher and a value of type []string which stores some of their accomplishments. Store 5x records in the map, printing out all values with index position in the slice

package main

import "fmt"


func main()  {
	bp := map[string][]string{
		`Turing, Alan`: []string{`Creator of Enigma "Prof's Book"`, `Developed statistical procedure 'Banburismus' to increase Bombe machine efficiency`, `Created 'Applications of Probability to Cryptography and 'Paper on Statistics of Repetitions'`},
		`Welchman, Gordon`: []string{`Member of the "Wicked Uncles"`, `Developed metadata traffic analysis process`, `Head of Hut 6, Enigma cipher break`},
		`Batey, Mavis`: []string{`Hut 6 codebreaker`, `Battle of Matapan decryption`, `Abwehr Enigma message decoding`},
		`Fawcett, Jane`:	[]string{`Hut 6 codebreaker`, `Decoded message detailing Bismarck position`, `Secretary, Victorian Society` },
		`Valentine, Jean`:	[]string{`Bombe operator`, `WREN`, `Bombe reconstruction, post war`},
	}
fmt.Println("DECLASSIFIED")
	for key, value := range bp {
		fmt.Println("Bletchley Personnel Record: ", key)
		for index, value := range value {
			fmt.Println("\t", index, value)
		}
	}
}
