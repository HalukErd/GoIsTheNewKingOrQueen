package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
}

func main() {
	t := triangle{base: 10, height: 20}
	s := square{length: 15}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func printArea(s shape) {
	fmt.Printf("Area: %v\n", s.getArea())
}
