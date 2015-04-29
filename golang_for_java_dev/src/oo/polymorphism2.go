package main

import (
	"fmt"
	"math"
)

type (
	Shape interface {
		Area() float64
		Perimeter() float64
	}

	Triangle struct {
		a, b, c float64
	}

	Rectangle struct {
		a, b float64
	}
)

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.a + r.b)
}

func main() {
	t := Triangle{4.0, 13.0, 15.0}
	r := Rectangle{4.0, 13.0}

	printShape(t)
	printShape(r)
}

func printShape(s Shape) {
	a := s.Area()
	p := s.Perimeter()

	fmt.Printf("%T's area and perimeter is %f and %f\n", s, a, p)
}

