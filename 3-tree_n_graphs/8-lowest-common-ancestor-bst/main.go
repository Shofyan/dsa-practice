package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor finds the LCA of two nodes in a BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// If both p and q are less than root, LCA is in left subtree
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// If both p and q are greater than root, LCA is in right subtree
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// Otherwise, root is the LCA
	return root
}

// Helper to insert a value into the BST
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

// Helper to find a node by value
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return findNode(root.Left, val)
	}
	return findNode(root.Right, val)
}

func main() {
	/*
			6
		   / \
		  2   8
		 / \ / \
		0  4 7  9
		  / \
		 3   5
	*/
	vals := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	var root *TreeNode
	for _, v := range vals {
		root = insert(root, v)
	}
	p := findNode(root, 2)
	q := findNode(root, 8)
	lca := lowestCommonAncestor(root, p, q)
	fmt.Printf("LCA of %d and %d is %d\n", p.Val, q.Val, lca.Val)

	p = findNode(root, 2)
	q = findNode(root, 4)
	lca = lowestCommonAncestor(root, p, q)
	fmt.Printf("LCA of %d and %d is %d\n", p.Val, q.Val, lca.Val)
}
