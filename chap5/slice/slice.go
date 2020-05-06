package main

import "fmt"

func main() {
	x := make([]float64, 5)
	y := x[0:2]

	fmt.Println(y)

	z := append(y, 4, 5)
	fmt.Println(z)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice2)
}
