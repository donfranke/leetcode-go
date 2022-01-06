package main

import "fmt"

func strStr(haystack string, needle string) int {

	if len(needle) < 1 {
		return 0
	} else {
		if len(haystack) < 1 || len(haystack) < len(needle) {
			return -1
		}
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if haystack == needle {
				return i
			}
			for j := 0; j < len(needle); j++ {
				//fmt.Printf("%d, %d, %s, %s\n", i, j, string(haystack[i+j]), string(needle[j]))
				if (i+j > len(haystack)-1 || i == len(haystack)-1) && j > 0 {
					return -1
				}
				if haystack[i+j] != needle[j] {
					break
				} else {
					if j == len(needle)-1 {
						return i
					}
				}
			}
		}
	}
	return -1
}

func main() {
	haystack := "abc"
	findStr := "c"

	ret := strStr(haystack, findStr)
	fmt.Println(ret)
}
