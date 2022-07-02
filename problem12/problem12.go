// https://leetcode.com/problems/search-in-a-binary-search-tree/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// 1. search binary tree

	return root
}

func main() {
	fmt.Println("Hello World")

	// build Btree for testing purposes only

	// Input: root = [4,2,7,1,3], val = 2
	nums := []int{4, 2, 7, 1, 3}
	// create binary tree
	var t TreeNode
	// root node
	t.Val = nums[0]

	// iterate remaining nums
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			t.Left = &TreeNode{Val: nums[i]}
		} else {
			t.Right = &TreeNode{Val: nums[i]}
		}
	}

	// Output: [2,1,3]

	// build binary tree

}
