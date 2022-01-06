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
	//ary := []int{1, 1, 2}
	//ary := []int{1, 2}
	//ary := []int{1, 2, 3}
	//ary := []int{1}
	ary := []int{}
	//ary := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(ary) // Calls your implementation
	fmt.Printf("Len: %d\n", k)
}
