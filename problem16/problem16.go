package main

import "fmt"

func main() {
	fmt.Println("Test it, don't run it")
}

func numIdenticalPairs(nums []int) int {
	pairCount := 0
	for i, j := range nums {
		for k, l := range nums {
			if i < k && j == l {
				pairCount++
			}
		}
	}
	return pairCount
}
