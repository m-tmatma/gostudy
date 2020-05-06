package main

import "fmt"

func main() {
	//var x map[string]int
	x := make(map[string]int)
	x["key"] = 10
	x["key2"] = 12

	fmt.Println(x["key"])

	for i, data := range x {
		fmt.Println(i, data)
	}

	if name, ok := x["key"]; ok {
		fmt.Println(name, ok)
	}
	if name, ok := x["key3"]; ok {
		fmt.Println(name, ok)
	}

	y := make(map[int]int)
	y[1] = 10
	y[10] = 20

	fmt.Println(y[1])

	for i, data := range y {
		fmt.Println(i, data)
	}
}
