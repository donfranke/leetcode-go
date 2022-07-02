package main

import (
	"testing"
)

func Test1(a *testing.T) {
	testUrl := "https://leetcode.com/problems/design-tinyurl"
	expected := testUrl

	obj := Constructor()
	encodedUrl := obj.encode(testUrl)
	decodedUrl := obj.decode(encodedUrl)

	if decodedUrl != expected {
		a.Errorf("Did not encode or decode correctly, result: %s\n", decodedUrl)
	}
}

func Test2(a *testing.T) {
	testUrl := "https://www.example.com/art"
	expected := testUrl

	obj := Constructor()
	encodedUrl := obj.encode(testUrl)
	decodedUrl := obj.decode(encodedUrl)

	if decodedUrl != expected {
		a.Errorf("Did not encode or decode correctly, result: %s\n", decodedUrl)
	}
}

func Test3(a *testing.T) {
	testUrl := "http://angle.example.com/"
	expected := testUrl

	obj := Constructor()
	encodedUrl := obj.encode(testUrl)
	decodedUrl := obj.decode(encodedUrl)

	if decodedUrl != expected {
		a.Errorf("Did not encode or decode correctly, result: %s\n", decodedUrl)
	}
}
