package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// validateBST checks if a binary tree is a valid BST
func validateBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

func main() {
	// Valid BST
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 2}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 8}

	// Invalid BST (right child of left subtree is greater than root)
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 7}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 6} // Invalid: 6 > 5 but in left subtree

	fmt.Println("root1 is valid BST:", validateBST(root1)) // true
	fmt.Println("root2 is valid BST:", validateBST(root2)) // false
}
