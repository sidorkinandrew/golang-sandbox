package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println(r.Area(), r.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())
}
