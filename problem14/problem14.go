package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test it, don't run it")
}

func countGoodRectangles(rectangles [][]int) int {
	m := make(map[int]int)
	for _, elem := range rectangles {
		// get smaller of 2 values
		minVal := 0
		if elem[0] < elem[1] {
			minVal = elem[0]
		} else {
			minVal = elem[1]
		}
		m[minVal]++
	}

	// return count of most popular measurement
	maxLen := 0
	for k, _ := range m {
		if k > maxLen {
			maxLen = k
		}
	}
	return m[maxLen]
}
