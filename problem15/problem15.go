// https://leetcode.com/problems/delete-node-in-a-linked-list/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// create linked inst
	// [4,5,1,9], node = 5 -> 4,1,9
	head := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}}
	deleteNode(head.Next)
	for head.Next != nil {
		fmt.Printf("%d\n", head.Val)
		head = head.Next
	}
	fmt.Printf("%d", head.Val)
}

func deleteNode(node *ListNode) {
	// shift node to delete to next node
	*node = *node.Next
}
