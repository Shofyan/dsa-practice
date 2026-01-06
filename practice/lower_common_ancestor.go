package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case
	if root == nil || root == p || root == q {
		return root
	}

	// Search left and right subtrees
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If p and q are found in different subtrees
	if left != nil && right != nil {
		return root
	}

	// Otherwise return the non-nil child
	if left != nil {
		return left
	}
	return right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
