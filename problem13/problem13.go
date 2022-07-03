// https://leetcode.com/problems/merge-two-binary-trees/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
		return root1
	}
	if root1 != nil {
		return root1
	} else {
		return root2
	}
}

func printTree(root *TreeNode) *TreeNode {
	if root != nil {
		fmt.Println(root.Val)
		root.Left = printTree(root.Left)
		root.Right = printTree(root.Right)
		return root
	}
	return root
}

func main() {

	// build bt1
	bt1 := &TreeNode{Val: 1}
	bt1.Left = &TreeNode{Val: 3}
	bt1.Right = &TreeNode{Val: 2}
	bt1.Left.Left = &TreeNode{Val: 5}

	// build bt2
	bt2 := &TreeNode{Val: 2}
	bt2.Left = &TreeNode{Val: 1}
	bt2.Right = &TreeNode{Val: 3}
	bt2.Left.Right = &TreeNode{Val: 4}
	bt2.Right.Right = &TreeNode{Val: 7}

	root3 := mergeTrees(bt1, bt2)
	printTree(root3)
}

// Insert node into tree - to build btree for testing purposes only
// Based on https://www.youtube.com/watch?v=-oYitelECuQ&t=491s
func (n *TreeNode) Insert(k int) {
	if k > 0 {
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
}
