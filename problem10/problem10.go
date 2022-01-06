// https://leetcode.com/problems/binary-search/

package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	// get middle number
	middle_number := nums[len(nums)/2]
	middle_index := len(nums) / 2

	// if number is in the middle of the array
	if middle_number == target {
		return middle_index
	}

	// if number is in first half of array
	if target < middle_number {
		nums = nums[:middle_index]
		for i, n := range nums {
			if n == target {
				return i
			}
		}
	}

	// if number is in second half of array
	if target > middle_number {
		index := len(nums[:middle_index])
		nums = nums[middle_index:]
		for i, n := range nums {
			if n == target {
				return index + i
			}
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 0
	ret := search(nums, target)
	fmt.Println(ret)

}
