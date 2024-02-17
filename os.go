package main

import "os"

func main() {

	args := os.Args
	for _, arg := range args {
		println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		println(hostname)
	} else {
		println("Error", err.Error())
	}

}
