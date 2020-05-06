package main

import "fmt"

func main() {
	x := make([]float64, 5)
	y := x[0:2]

	fmt.Println(y)

	z := append(y, 4, 5)
	fmt.Println(z)
}
