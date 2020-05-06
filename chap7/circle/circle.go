package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	c := new(Circle)
	d := Circle{1, 2, 3}
	e := Circle{x: 1, y: 2, r: 3}

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
