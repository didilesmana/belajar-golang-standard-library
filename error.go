package main

import (
	"errors"
)

var (
	ValidationError = errors.New("valid error")
	NotFoundError   = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "didi" {
		return NotFoundError
	}

	return nil

}

func main() {
	err := GetById("didi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("valid eror")
		} else if errors.Is(err, NotFoundError) {
			println("not found")
		} else {
			println("unknown")
		}
	} else {
		println("sukses")
	}

}
