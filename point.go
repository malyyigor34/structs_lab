package main

import "fmt"

type point struct {
	a, b float64
}

func (p point) String() string {
	return fmt.Sprintf("(%f;%f)", p.a, p.b)
}
