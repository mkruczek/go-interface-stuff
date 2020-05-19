package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("%f\n", s.getArea())
}

type triangle struct {
	h float64
	a float64
}
type square struct {
	a float64
}

func (t triangle) getArea() float64 {
	return t.h * 0.5 * t.a
}

func (s square) getArea() float64 {
	return s.a * s.a
}

func main() {
	t := triangle{h: 3, a: 5}
	s := square{a: 5}

	printArea(t)
	printArea(s)

}
