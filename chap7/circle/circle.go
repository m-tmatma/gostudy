package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	c := new(Circle)

	fmt.Println(c)
}
