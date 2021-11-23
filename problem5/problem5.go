package main

import "fmt"

func main() {

	var strs = []string{"ab", "a"}

	str_result := longestCommonPrefix(strs)
	fmt.Println("PREFIX: ", str_result)
}

func longestCommonPrefix(strs []string) string {
	max_length := 200
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < max_length {
			max_length = len(strs[i])
		}
	}

	prefix := ""
	for j := 0; j < max_length; j++ {
		first_letter := strs[0][j : j+1]

		for i := 1; i < len(strs); i++ {

			if first_letter != strs[i][j:j+1] {
				return prefix
			}
		}
		prefix += first_letter
	}
	return prefix
}
