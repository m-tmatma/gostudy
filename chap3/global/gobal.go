package main

import "fmt"

var x string = "Hello"

func main() {
	fmt.Println(x)
	f()
}

func f() {
	fmt.Println(x)
}
