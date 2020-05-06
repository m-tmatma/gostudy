package main

import (
	"fmt"
	"math"
)

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var min = math.MaxInt64
	for _, val := range x {
		if min > val {
			min = val
		}
	}
	fmt.Println(min)
	fmt.Println(x)
}
