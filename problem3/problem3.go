package main

import (
	"fmt"
	"strconv"
)

func main() {
	b_ret := isPalindrome(121)
	fmt.Println(b_ret)
}

func isPalindrome(x int) bool {
	new_str := ""
	x_str := strconv.Itoa(x)
	for i := len(x_str); i > 0; i-- {
		new_str += x_str[i-1 : i]
		if x_str == new_str {
			return true
		}
	}
	return false
}
