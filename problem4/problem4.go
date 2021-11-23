package main

import (
	"fmt"
	"os"
)

func main() {

	roman_number_number := os.Args[1]

	i_ret := romanToInt(roman_number_number)
	fmt.Println("Return: ", i_ret)
}

func romanToInt(s string) int {
	integer_value := 0

	roman_numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	i_total := 0

	for i := 0; i < len(s); i++ {
		curr_letter := s[i : i+1]
		curr_value := roman_numerals[curr_letter]

		if i < len(s)-1 {
			if roman_numerals[s[i+1:i+2]] > curr_value {
				i_total = roman_numerals[s[i+1:i+2]] - curr_value
				i++
			} else {
				i_total = curr_value
			}

		} else {
			i_total = roman_numerals[s[i:i+1]]
		}
		integer_value += i_total
	}
	return integer_value
}
