package main

import "fmt"

func main() {
	var x string = "Hello, World"

	fmt.Println(x)

	var y string
	y = "Hello"
	fmt.Println(y)
	y = x + "Second"
	fmt.Println(y)
}
