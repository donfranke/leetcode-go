package main

import "testing"

func Test1(a *testing.T) {
	if isValid("A") == false {
		a.Error("A should be true")
	}
}
func Test2(a *testing.T) {
	if isValid("B") == true {
		a.Error("B should be false")
	}
}
