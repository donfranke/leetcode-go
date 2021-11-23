package main

import "testing"

func Test1(a *testing.T) {
	if isValid("()") == false {
		a.Error("() should be true")
	}
}
func Test2(a *testing.T) {
	if isValid("[({(())}[()])]") == false {
		a.Error("[({(())}[()])] should be true")
	}
}
func Test3(a *testing.T) {
	if isValid("(())[()]") == false {
		a.Error("(())[()] should be true")
	}
}

func Test4(a *testing.T) {
	if isValid("()[]{}") == false {
		a.Error("()[]{} should be true")
	}
}

func Test5(a *testing.T) {
	if isValid("(]") == true {
		a.Error("(] should be false")
	}
}
func Test6(a *testing.T) {
	if isValid("([)]") == true {
		a.Error("([)] should be false")
	}
}
func Test7(a *testing.T) {
	if isValid("{[]}") == false {
		a.Error("{[]} should be true")
	}
}
func Test8(a *testing.T) {
	if isValid("(){}}{") == true {
		a.Error("(){}}{ should be false")
	}
}
func Test9(a *testing.T) {
	if isValid("{}}{") == true {
		a.Error("{}}{ should be false")
	}
}
func Test10(a *testing.T) {
	if isValid("([}}])") == true {
		a.Error("([}}]) should be false")
	}
}
func Test11(a *testing.T) {
	if isValid("(([]){})") == false {
		a.Error("(([]){}) should be true")
	}
}

func Test12(a *testing.T) {
	if isValid("((") == true {
		a.Error("(( should be false")
	}
}
