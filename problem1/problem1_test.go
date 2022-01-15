package main

import (
	"testing"
)

func Test1(a *testing.T) {
	aryInput := []int{1, 2, 1}
	aryExpected := []int{1, 2, 1, 1, 2, 1}
	aryResults := getConcatenation(aryInput)
	// compare results
	for i, v1 := range aryResults {
		if aryExpected[i] != v1 {
			a.Error("Mismatched arrays")
		}
	}
}

func Test2(a *testing.T) {
	aryInput := []int{1, 3, 2, 1}
	aryExpected := []int{1, 3, 2, 1, 1, 3, 2, 1}
	aryResults := getConcatenation(aryInput)
	// compare results
	for i, v1 := range aryResults {
		if aryExpected[i] != v1 {
			a.Error("Mismatched arrays")
		}
	}
}
