// https://leetcode.com/problems/root-equals-sum-of-children/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Left.Val+root.Right.Val == root.Val
}

func main() {
	// Input: root = [5,3,1]
	// Output: false
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 1}
	fmt.Println(checkTree(root))
}

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
