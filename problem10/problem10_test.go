package main

import (
	"testing"
)

func Test1(a *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 0

	if search(nums, target) != 1 {
		a.Error("Answer should be 1")
	}
}

func Test2(a *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 5

	if search(nums, target) != 3 {
		a.Error("Answer should be 3")
	}
}

func Test3(a *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 11

	if search(nums, target) != -1 {
		a.Error("Answer should be -1")
	}
}
