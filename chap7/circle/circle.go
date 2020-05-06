package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := new(Circle)
	d := Circle{1, 2, 3}
	e := Circle{x: 1, y: 2, r: 3}

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(circleArea(e))
	fmt.Println(circleArea2(&e))
	fmt.Println(e.area())
}
