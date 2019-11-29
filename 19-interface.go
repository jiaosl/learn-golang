package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//rect
func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perimeter() float64 {
	return r.height*2 + r.width*2
}

// circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g", g)
	fmt.Println("g.area()", g.area())
	fmt.Println("g.perimeter():", g.perimeter())
}
func main() {
	r := rect{3, 4}
	c := circle{5}
	measure(r)
	measure(c)
}
