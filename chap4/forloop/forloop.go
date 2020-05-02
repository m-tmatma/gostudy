package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	for k := 1; k <= 10; k++ {
		switch k {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown")
		}
	}
}
