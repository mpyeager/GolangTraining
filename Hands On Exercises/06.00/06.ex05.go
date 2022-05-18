// Level 6: Exercise 5
// Create a type square and type circle. Attach a method to each which calculates and returns area. Create a type shape with defines an interface as anything which has the area method. Create a function info which takes type shape and then prints the area. Create a value of type squre and a value of type circle. Use func info to print the areas of the square and circle.

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape)  {
	fmt.Println(s.area())
}
func main()  {
	c := circle{
		radius: 43.75,
	}
	sq := square{
		length: 18,
	}

	info(c)
	info(sq)
}
