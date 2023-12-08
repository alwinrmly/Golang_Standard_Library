package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) len() int {
	return len(s)
}

func (s UserSlice) less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Alwin", 26},
		{"Darsono", 30},
		{"Asep", 29},
		{"Tejo", 35},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
