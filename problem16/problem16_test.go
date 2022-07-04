package main

import (
	"testing"
)

func Test1(a *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	expected := 4
	returned := numIdenticalPairs(nums)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test2(a *testing.T) {
	nums := []int{1, 1, 1, 1}
	expected := 6
	returned := numIdenticalPairs(nums)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test3(a *testing.T) {
	nums := []int{1, 2, 3}
	expected := 0
	returned := numIdenticalPairs(nums)

	if returned != expected {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}
