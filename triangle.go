package main

import (
	"fmt"
	"math"
)

type triangle struct {
	a, b, c line
}

func (t triangle) perimeter() float64 {
	return t.a.len() + t.b.len() + t.c.len()
}

func (t triangle) String() string {
	return fmt.Sprintf("Triangle. length: %f, %f, %f, perimeter: %f, square: %f", t.a.len(), t.b.len(), t.c.len(), t.perimeter(),
		t.square())
}

func (t triangle) square() float64 {
	return math.Sqrt(0.5 * t.perimeter() * (0.5*t.perimeter() - t.a.len()) * (0.5*t.perimeter() - t.b.len()) * (0.5*t.perimeter() - t.c.len()))
}
