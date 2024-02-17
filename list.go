package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Didi")
	data.PushBack("Lesmana")
	data.PushBack(100)

	head := data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
