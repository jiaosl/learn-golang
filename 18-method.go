package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}
func (r *rect) perimeter() int {
	return r.width*2 + r.height*2
}

func main() {
	r := rect{10, 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
