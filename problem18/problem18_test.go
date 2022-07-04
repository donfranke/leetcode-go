package main

import (
	"testing"
)

func Test1(a *testing.T) {
	s := "ababcbacadefegdehijhklij"
	expected := []int{9, 7, 8}
	returned := partitionLabels(s)
	testPassed := true
	for k, v := range returned {
		if v != expected[k] {
			testPassed = false
		}
	}

	if testPassed == false {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test2(a *testing.T) {
	s := "eccbbbbdec"
	expected := []int{10}
	returned := partitionLabels(s)
	testPassed := true
	for k, v := range returned {
		if v != expected[k] {
			testPassed = false
		}
	}

	if testPassed == false {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}

func Test3(a *testing.T) {
	s := "caedbdedda"
	expected := []int{1, 9}
	returned := partitionLabels(s)
	testPassed := true
	for k, v := range returned {
		if v != expected[k] {
			testPassed = false
		}
	}

	if testPassed == false {
		a.Errorf("RESULTS --> Expected: %d, Returned: %d\n", expected, returned)
	}
}
