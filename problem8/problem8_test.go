package main

import (
	"testing"
)

func Test1(a *testing.T) {
	testAry := []int{1, 1, 2}

	if removeDuplicates(testAry) != 2 {
		a.Error("Answer is 2")
	}

}

func Test2(a *testing.T) {
	testAry := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	if removeDuplicates(testAry) != 5 {
		a.Error("Answer is 5")
	}

}
