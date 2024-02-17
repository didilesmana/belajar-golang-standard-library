package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Umur int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Umur < s[j].Umur
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"didi", 25},
		{"budi", 30},
		{"adit", 20},
		{"yudi", 45},
		{"nana", 18},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}