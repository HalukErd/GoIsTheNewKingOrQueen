package main

import (
	"fmt"
	"math"
	"sync"
)

type rectangle struct {
	a, b float64
}
type circle struct {
	r float64
}
type triangle struct {
	a, h float64
}
type shape interface {
	area() float64
	printInfo()
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	r := rectangle{23, 54}
	c := circle{12}
	t := triangle{9, 10}
	go printShapeInfo(r)
	go printShapeInfo(c)
	go printShapeInfo(t)
	wg.Wait()
	//printArea(r)
	//printArea(c)
	//printArea(t)
}
func printArea(s shape) {
	fmt.Printf("The %T's area: %.2f\n", s, s.area())
}
func printShapeInfo(s shape) {
	s.printInfo()
}
func (r rectangle) area() float64 {
	return r.a * r.b
}
func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (t triangle) area() float64 {
	return t.h * t.a / 2
}
func (r rectangle) printInfo() {
	fmt.Printf("Rectangle's width and height: %v, %v --- Rectangle's Area:%v\n", r.a, r.b, r.area())
	wg.Done()
}
func (c circle) printInfo() {
	fmt.Printf("Circle's radius: %v --- Circle's Area:%v\n", c.r, c.area())
	wg.Done()
}
func (t triangle) printInfo() {
	fmt.Printf("Triangle's side and side's height: %v, %v --- Triangle's Area:%v\n", t.a, t.h, t.area())
	wg.Done()
}
