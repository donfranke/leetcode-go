package main

import (
	"testing"
)

func Test1(a *testing.T) {

	// rectangles = [[5,8],[3,9],[5,12],[16,5]] -> 3
	rectangles := [][]int{
		{5, 8},
		{3, 9},
		{5, 12},
		{16, 5},
	}
	//fmt.Println(rectangles)
	expected := 3
	returned := countGoodRectangles(rectangles)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test2(a *testing.T) {

	// rectangles = [[2,3],[3,7],[4,3],[3,7]] -> 3
	rectangles := [][]int{
		{2, 3},
		{3, 7},
		{4, 3},
		{3, 7},
	}
	//fmt.Println(rectangles)
	expected := 3
	returned := countGoodRectangles(rectangles)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test3(a *testing.T) {

	// rectangles = [[5,8],[3,9],[3,12]] -> 1
	rectangles := [][]int{
		{5, 8},
		{3, 9},
		{3, 12},
	}
	//fmt.Println(rectangles)
	expected := 1
	returned := countGoodRectangles(rectangles)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}
