package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Rahmad"}
	values := []int{100, 90, 80, 85}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Alwin"))
	fmt.Println(slices.Index(names, "Alwin"))
	fmt.Println(slices.Index(names, "George"))
}
