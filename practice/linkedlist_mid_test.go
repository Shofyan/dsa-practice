package main

import "testing"

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode mengembalikan node tengah dari linked list.
//
// Parameter:
//   - head: pointer ke node pertama dari linked list
//
// Return:
//   - pointer ke node tengah dari linked list
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {

		println("slow:", slow.Val, "fast:", fast.Val)

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}

func Test_mid_linkedlist(t *testing.T) {
	// Membuat linked list: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

	mid := middleNode(head)
	if mid.Val != 4 {
		t.Errorf("Expected middle node value to be 4, but got %d", mid.Val)
	}
}
