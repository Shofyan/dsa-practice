package main

import "fmt"

// Node represents a node in the binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// BinarySearchTree represents the binary search tree
type BinarySearchTree struct {
	Root *Node
}

// Insert inserts a new key into the binary search tree
func (bst *BinarySearchTree) Insert(key int) {
	bst.Root = insert(bst.Root, key)
}

// insert is a recursive helper function to insert a key
func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key}
	}
	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	}
	return root
}

// Search searches for a key in the binary search tree
func (bst *BinarySearchTree) Search(key int) bool {
	return search(bst.Root, key)
}

// search is a recursive helper function to search for a key
func search(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if key < root.Key {
		return search(root.Left, key)
	} else if key > root.Key {
		return search(root.Right, key)
	}
	return true
}

// InOrder performs an in-order traversal of the tree
func (bst *BinarySearchTree) InOrder() {
	inOrder(bst.Root)
	fmt.Println()
}

// inOrder is a recursive helper function for in-order traversal
func inOrder(root *Node) {
	if root != nil {
		inOrder(root.Left)
		fmt.Printf("%d ", root.Key)
		inOrder(root.Right)
	}
}

func main() {
	bst := &BinarySearchTree{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(70)
	bst.Insert(60)
	bst.Insert(80)

	fmt.Print("In-order traversal: ")
	bst.InOrder()

	fmt.Println("Search for 40:", bst.Search(40))
	fmt.Println("Search for 90:", bst.Search(90))
}