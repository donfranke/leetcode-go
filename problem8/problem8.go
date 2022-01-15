// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

import "fmt"

func removeDuplicates(nums []int) int {
	total := 0

	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			if nums[i] == nums[len(nums)-1] {
				total = i
				break
			}
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			removeDuplicates(nums)
		}
		total++
	}
	return total + 1
}

func main() {
	ary := []int{}
	k := removeDuplicates(ary)
	fmt.Printf("Len: %d\n", k)
}
