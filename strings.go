package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Alwin Ramli", "Alwin"))
	fmt.Println(strings.Split("Alwin Ramli", " "))
	fmt.Println(strings.ToLower("Alwin Ramli"))
	fmt.Println(strings.ToUpper("Alwin Ramli"))
	fmt.Println(strings.Trim("         Alwin Ramli   ", " "))
	fmt.Println(strings.ReplaceAll("Alwin Ramli Darsono", "Darsono", "Sasmita"))
}
