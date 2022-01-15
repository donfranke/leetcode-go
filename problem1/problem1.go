// https://leetcode.com/problems/concatenation-of-array/

package main

import "fmt"

func main() {
	var aryInput = []int{1, 2, 1}
	var aryOutput []int

	aryOutput = getConcatenation(aryInput)
	for _, v := range aryOutput {
		fmt.Println(v)
	}
}

func getConcatenation(nums []int) []int {
	var ary = nums
	for _, v := range ary {
		ary = append(ary, v)
	}
	return ary
}
