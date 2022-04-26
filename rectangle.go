package main

import "fmt"

type rectangle struct {
	a, b, c, d line
}

func (r rectangle) square() float64 {
	return r.a.len() * r.b.len()
}

func (r rectangle) perimeter() float64 {
	return 2*r.a.len() + 2*r.b.len()
}

func (r rectangle) String() string {
	return fmt.Sprintf("Rectangle. lengths: %f %f %f %f; perimeter: %f, square: %f", r.a.len(), r.b.len(), r.c.len(), r.d.len(),
		r.perimeter(), r.square())
}
