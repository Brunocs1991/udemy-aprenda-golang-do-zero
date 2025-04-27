package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func writeArea(s shape) {
	fmt.Printf("The shape form is %0.2f: \n", s.area())
}

type rectangle struct {
	length, width float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

type trapezium struct {
	base1, base2, height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t trapezium) area() float64 {
	return 0.5 * (t.base1 + t.base2) * t.height
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{length: 10, width: 15}
	writeArea(r)

	c := circle{radius: 10}
	writeArea(c)

	triangle1 := triangle{base: 10, height: 15}
	writeArea(triangle1)

	trapezium1 := trapezium{base1: 10, base2: 5, height: 7}
	writeArea(trapezium1)
}
