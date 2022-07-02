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

// search binary tree
func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if val == root.Val {
			return root
		} else if val > root.Val {
			return searchBST(root.Right, val)
		} else if val < root.Val {
			return searchBST(root.Left, val)
		}
	}
	return nil
}

func main() {
	// build Btree for testing purposes only
	nums := []int{4, 2, 7, 1, 3}

	// create binary tree
	bt := &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		bt.Insert(nums[i])
	}
	fmt.Printf("Test: %v\n", bt)

	// search for root matching value
	bt_root := searchBST(bt, 5)
	fmt.Printf("Root: %v\n", bt_root)
}

// Insert node into tree - to build btree for testing purposes only
// Based on https://www.youtube.com/watch?v=-oYitelECuQ&t=491s
func (n *TreeNode) Insert(k int) {
	if k > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: k}
		} else {
			n.Right.Insert(k)
		}
	} else {
		if n.Left == nil {
			n.Left = &TreeNode{Val: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
