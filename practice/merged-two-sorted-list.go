package main

import "fmt"

// buat linkedlist
type listNode struct {
	Val  int
	Next *listNode
}

func MergeListNode(list1 *listNode, list2 *listNode) *listNode {
	dummyNode := &listNode{}
	current := dummyNode

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummyNode.Next

}

// ===== Helpers =====

// Build linked list from slice
func buildList(arr []int) *listNode {
	if len(arr) == 0 {
		return nil
	}
	head := &listNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &listNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

// Print linked list
func printList(head *listNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
