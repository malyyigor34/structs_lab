package main

import (
	"fmt"
	"math"
)

type line struct {
	a, b point
}

func (l line) String() string {
	return fmt.Sprintf("Line length: %f", l.len())
}

func (l line) len() float64 {
	return math.Sqrt(math.Pow(l.a.a-l.a.b, 2) + math.Pow(l.b.a-l.b.b, 2))
}
