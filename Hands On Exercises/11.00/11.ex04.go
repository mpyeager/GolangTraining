// Level 11: Exercise 4
// Building on the code below, use the sqrt.Error struct as a value of type error. Use the following lat "50.2289 N", long "99.4656 W"

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat 	string
	long 	string
	err		error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Norgate Math Errrror: %v %v %v", se.lat, se.long, se.err)
}

func main()  {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error)  {
	if f < 0 {
		e := errors.New("Earl Grey not detected. Moar tea required.")
		return 0, sqrtError {
			"50.2289 N",
			"99.4656 W",
			e,
		}
	return 42, nil
	}
}
