package main

import "errors"

func main() {
	println("Hello, World!")
	isValid("hello")
}

func isValid(txt string) error {
		if txt == "" {
		return errors.New("Invalid")
	}
	return nil
}
