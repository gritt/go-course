package main

import "fmt"

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {

	s := square{
		side: 2.2,
	}

	t := triangle{
		height: 1.0,
		base:   2.5,
	}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("the shape area is:", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}
