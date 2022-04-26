package main

import (
	"fmt"
	"math"
)

type circle struct {
	// a - średnica
	a line
}

func (c circle) square() float64 {
	return math.Pi * math.Pow(c.a.len()/2, 2)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.a.len() * 2
}

func (c circle) String() string {
	return fmt.Sprintf("Circle. Radius %f, perimeter: %f, square: %f", c.a.len()/2, c.perimeter(), c.square())
}
