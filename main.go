package main

import (
	"fmt"
)

func main() {
	create_circle()
	create_triangle()
	create_rectangle()
}

func create_rectangle() {
	l1 := line{point{1, 3}, point{3, 5}}
	l2 := line{point{1, 5}, point{4, 1}}
	l3 := line{point{4, 5}, point{1, 5}}
	l4 := line{point{4, 5}, point{1, 5}}
	r := rectangle{l1, l2, l3, l4}
	fmt.Println(r)
}

func create_circle() {
	l1 := line{point{1, 3}, point{3, 5}}
	c := circle{l1}
	fmt.Println(c)
}

func create_triangle() {
	l1 := line{point{1, 3}, point{3, 5}}
	l2 := line{point{1, 5}, point{4, 1}}
	l3 := line{point{4, 5}, point{1, 5}}
	t := triangle{l1, l2, l3}
	fmt.Println(t)
}
