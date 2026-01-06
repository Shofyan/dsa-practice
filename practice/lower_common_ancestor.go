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

func lowestCommonAncestorNorecursion(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)

	stack := []*TreeNode{root}
	parent[root] = nil

	// Traverse until both nodes are found
	for parent[p] == nil || parent[q] == nil {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// Record ancestors of p
	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}

	// Find first common ancestor of q
	for !ancestors[q] {
		q = parent[q]
	}

	return q
}
