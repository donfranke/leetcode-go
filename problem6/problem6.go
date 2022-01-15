// https://leetcode.com/problems/valid-parentheses/
package main

import (
	"fmt"
	"strings"
)

func main() {
	//parens_string := "[({(())}[()])]" // true
	parens_string := "((" // true
	b_result := isValid(parens_string)
	fmt.Println("IS VALID: ", b_result)
}

func isValid(s string) bool {
	symbol_pairs := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	if len(s)%2 > 0 {
		return false
	} else {
		for i := 0; i < len(s); i++ {
			this_char := s[i : i+1]
			map_char := symbol_pairs[this_char]
			if map_char == "" {
				return false
			}

			this_char_index := i
			map_char_index := strings.Index(s, map_char)

			if map_char_index < 1 {
				return false
			}

			if map_char_index-this_char_index == 1 {

				if len(s) == 2 {
					return true
				} else {
					left_string := s[:this_char_index]
					right_string := s[map_char_index+1:]
					new_string := left_string + right_string
					return isValid(new_string)
				}
			}
		}
	}
	return true
}
