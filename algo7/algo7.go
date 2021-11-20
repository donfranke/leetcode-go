package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	ret := isValid("B")
	fmt.Println(ret)
}

func isValid(s string) bool {
	return s == "A"
}
