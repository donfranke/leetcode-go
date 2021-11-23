package main

import "fmt"

func main() {
	var ary_nums = []int{-1, 0, 3, 5, 9, 12}
	var iTarget = 9

	var i_result int
	i_result = search(ary_nums, iTarget)
	fmt.Println(i_result)
}

func search(nums []int, target int) int {
	i := -1
	for i = 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
