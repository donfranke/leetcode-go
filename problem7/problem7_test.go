package main

import (
	"fmt"
	"testing"
)

func Test1(a *testing.T) {

	//Input: l1 = [1,2,4], l2 = [1,3,4]
	//Output: [1,1,2,3,4,4]

	LN1 := ListNode{Val: 1}
	LN1.Next = &ListNode{Val: 2}
	LN1.Next.Next = &ListNode{Val: 4}

	LN2 := ListNode{Val: 1}
	LN2.Next = &ListNode{Val: 3}
	LN2.Next.Next = &ListNode{Val: 4}

	LNReturn := mergeTwoLists(&LN1, &LN2)

	temp := LNReturn
	i := 0
	for temp != nil {
		if i == 0 && temp.Val != 1 {
			a.Error("Node 1 value should be 1")
		}
		if i == 1 && temp.Val != 1 {
			a.Error("Node 2 value should be 1")
		}
		if i == 2 && temp.Val != 2 {
			a.Error("Node 3 value should be 2")
		}
		if i == 3 && temp.Val != 3 {
			a.Error("Node 4 value should be 3")
		}
		if i == 4 && temp.Val != 4 {
			a.Error("Node 5 value should be 4")
		}
		if i == 5 && temp.Val != 4 {
			a.Error("Node 6 value should be 4")
		}
		temp = temp.Next
		i += 1
	}
}

func Test2(a *testing.T) {

	// Input: l1 = [], l2 = []
	// Output: []

	LN1 := ListNode{}
	LN2 := ListNode{}

	LNReturn := mergeTwoLists(&LN1, &LN2)

	temp := LNReturn
	i := 0
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		if i == 0 && temp.Val != 0 {
			a.Error("Node 1 value should be 0")
		}
		temp = temp.Next
		i += 1
	}
}

func Test3(a *testing.T) {

	// Input: l1 = [], l2 = [0]
	// Output: [0]

	LN1 := ListNode{}
	LN2 := ListNode{Val: 0}

	LNReturn := mergeTwoLists(&LN1, &LN2)

	temp := LNReturn
	i := 0
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		if i == 0 && temp.Val != 0 {
			a.Error("Node 1 value should be 0")
		}
		temp = temp.Next
		i += 1
	}
}

func Test4(a *testing.T) {

	//Input: l1 = [1,2,4], l2 = [1,3,4]
	//Output: [1,1,2,3,4,4]

	LN1 := ListNode{Val: 107}
	LN1.Next = &ListNode{Val: 282}
	LN1.Next.Next = &ListNode{Val: 490}

	LN2 := ListNode{Val: 140}
	LN2.Next = &ListNode{Val: 301}
	LN2.Next.Next = &ListNode{Val: 499}

	LNReturn := mergeTwoLists(&LN1, &LN2)

	temp := LNReturn
	i := 0
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		if i == 0 && temp.Val != 107 {
			a.Error("Node 1 value should be 107")
		}
		if i == 1 && temp.Val != 140 {
			a.Error("Node 2 value should be 140")
		}
		if i == 2 && temp.Val != 282 {
			a.Error("Node 3 value should be 282")
		}
		if i == 3 && temp.Val != 301 {
			a.Error("Node 4 value should be 301")
		}
		if i == 4 && temp.Val != 490 {
			a.Error("Node 5 value should be 490")
		}
		if i == 5 && temp.Val != 499 {
			a.Error("Node 6 value should be 499")
		}
		temp = temp.Next
		i += 1
	}
}
