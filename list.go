package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Alwin")
	data.PushBack("Ramli")
	data.PushBack("Sasmita")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // return data alwin

	next := head.Next()
	fmt.Println(next.Value) // return data ramli

	next1 := next.Next()
	fmt.Println(next1.Value) // return data sasmita

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
