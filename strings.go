package main

import "strings"

func main() {
	println(strings.Contains("Didi Lesmana", "Didi"))
	println(strings.Split("Didi Lesmana", " "))
	println(strings.ToLower("Didi Lesmana"))
	println(strings.ToUpper("Didi Lesmana"))
	println(strings.Trim("     Didi Lesmana     ", " "))
	println(strings.ReplaceAll("Didi Lesmana Didi", "Didi", "Budi"))
}
