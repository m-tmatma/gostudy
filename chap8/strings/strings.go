package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "t"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
}
