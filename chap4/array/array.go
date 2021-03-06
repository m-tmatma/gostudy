package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 84

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	total = 0
	for _, value := range y {
		total += value
	}
	fmt.Print(total / float64(len(y)))
}
