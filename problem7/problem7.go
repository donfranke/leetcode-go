package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	cur_node := &ListNode{}
	all_node := cur_node

	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			cur_node.Next = l1
			l1 = l1.Next
		} else if l2 != nil && (l1 == nil || l1.Val >= l2.Val) {
			cur_node.Next = l2
			l2 = l2.Next
		}
		cur_node = cur_node.Next
	}
	return all_node.Next

}

func main() {
	fmt.Println("Use test script")

	LN1 := ListNode{Val: 1}
	LN1.Next = &ListNode{Val: 2}
	LN1.Next.Next = &ListNode{Val: 4}

	LN2 := ListNode{Val: 1}
	LN2.Next = &ListNode{Val: 3}
	LN2.Next.Next = &ListNode{Val: 4}

	ret_list := mergeTwoLists(&LN1, &LN2)
	current := ret_list
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}

}
