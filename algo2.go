package main

import "fmt"

func main() {
	var ary_nums = []int{2, 7, 11, 15}
	var iTarget = 9

	var intArr []int

	intArr = twoSum(ary_nums, iTarget)
	fmt.Printf("[%d, %d]", intArr[0], intArr[1])
}

func twoSum(nums []int, target int) []int {
	i := -1
	j := -1
	var intArr = []int{-1, -1}

	for i = 0; i < len(nums); i++ {
		for j = 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				intArr[0] = i
				intArr[1] = j
				return intArr
			}
		}
	}
	return intArr
}
