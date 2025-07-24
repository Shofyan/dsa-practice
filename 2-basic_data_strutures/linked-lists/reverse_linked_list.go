package main

import "fmt"

// ListNode defines a node in the linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList reverses a singly linked list.
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// printList prints the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Create a sample linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Print("Original list: ")
	printList(head)

	reversedHead := reverseList(head)

	fmt.Print("Reversed list: ")
	printList(reversedHead)
}
