package main

import (
	"fmt"
	"math"
)

func plus(a int, b int) int {
	return a + b
}

func circle(a int) (int, int, int) {
	b := a * 2
	c := a + 2
	d := math.Sin(float64(a))
	return b, c, int(d)
}
func main() {
	res := plus(1, 3)
	fmt.Println("res:", res)
	res1, _, res3 := circle(3)
	fmt.Println("res1, res3:", res1, res3)
}
