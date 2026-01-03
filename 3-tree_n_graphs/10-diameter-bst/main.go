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

// diameterOfBinaryTree returns the diameter of the binary tree
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		diameter = int(math.Max(float64(diameter), float64(left+right)))
		return 1 + int(math.Max(float64(left), float64(right)))
	}
	depth(root)
	return diameter
}

func main() {
	/*
				1
			   / \
			  2   3
			 / \
			4   5
		The diameter is 3 (path: 4 -> 2 -> 1 -> 3 or 5 -> 2 -> 1 -> 3)
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println("Diameter of the tree:", diameterOfBinaryTree(root)) // Output: 3

	// Test with single node
	single := &TreeNode{Val: 10}
	fmt.Println("Diameter of single node tree:", diameterOfBinaryTree(single)) // Output: 0
}
