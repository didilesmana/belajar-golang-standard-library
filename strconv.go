package main

import "strconv"

func main() {
	result, err := strconv.ParseBool("false")
	if err != nil {
		println("Error", err.Error())
	} else {
		println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		println("Error", err.Error())
	} else {
		println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	println(binary)

	var stringInt string = strconv.Itoa(999)
	println(stringInt)

}
