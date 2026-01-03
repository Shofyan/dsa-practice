package main

import (
	"fmt"
)

// SegmentTree represents a segment tree for range sum queries
type SegmentTree struct {
	tree []int
	n    int
}

// NewSegmentTree creates a new segment tree from the given array
func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.build(arr, 0, 0, n-1)
	return st
}

// build constructs the segment tree recursively
func (st *SegmentTree) build(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
	} else {
		mid := (start + end) / 2
		st.build(arr, 2*node+1, start, mid)
		st.build(arr, 2*node+2, mid+1, end)
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

// Query returns the sum of elements in range [l, r]
func (st *SegmentTree) Query(l, r int) int {
	return st.query(0, 0, st.n-1, l, r)
}

// query is the recursive helper for range sum queries
func (st *SegmentTree) query(node, start, end, l, r int) int {
	if r < start || end < l {
		return 0 // no overlap
	}
	if l <= start && end <= r {
		return st.tree[node] // complete overlap
	}

	// partial overlap
	mid := (start + end) / 2
	leftSum := st.query(2*node+1, start, mid, l, r)
	rightSum := st.query(2*node+2, mid+1, end, l, r)
	return leftSum + rightSum
}

// Update changes the value at index idx to val
func (st *SegmentTree) Update(idx, val int) {
	st.update(0, 0, st.n-1, idx, val)
}

// update is the recursive helper for point updates
func (st *SegmentTree) update(node, start, end, idx, val int) {
	if start == end {
		st.tree[node] = val
	} else {
		mid := (start + end) / 2
		if idx <= mid {
			st.update(2*node+1, start, mid, idx, val)
		} else {
			st.update(2*node+2, mid+1, end, idx, val)
		}
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

// Display prints the segment tree structure
func (st *SegmentTree) Display() {
	fmt.Println("Segment Tree structure:")
	for i := 0; i < len(st.tree) && st.tree[i] != 0; i++ {
		fmt.Printf("tree[%d] = %d\n", i, st.tree[i])
	}
	fmt.Println()
}

func main() {
	// Test array
	arr := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("Original array: %v\n", arr)

	// Create segment tree
	st := NewSegmentTree(arr)
	st.Display()

	// Test range queries
	fmt.Println("Range Query Tests:")
	fmt.Printf("Sum of range [1, 3]: %d\n", st.Query(1, 3)) // 3+5+7 = 15
	fmt.Printf("Sum of range [0, 2]: %d\n", st.Query(0, 2)) // 1+3+5 = 9
	fmt.Printf("Sum of range [2, 5]: %d\n", st.Query(2, 5)) // 5+7+9+11 = 32
	fmt.Printf("Sum of range [0, 5]: %d\n", st.Query(0, 5)) // 1+3+5+7+9+11 = 36
	fmt.Printf("Sum of range [4, 4]: %d\n", st.Query(4, 4)) // 9
	fmt.Println()

	// Test updates
	fmt.Println("Update Tests:")
	fmt.Printf("Before update - Sum of range [1, 3]: %d\n", st.Query(1, 3))

	st.Update(2, 10)                                                                       // Change arr[2] from 5 to 10
	fmt.Printf("After updating index 2 to 10 - Sum of range [1, 3]: %d\n", st.Query(1, 3)) // 3+10+7 = 20

	st.Update(0, 2)                                                                       // Change arr[0] from 1 to 2
	fmt.Printf("After updating index 0 to 2 - Sum of range [0, 2]: %d\n", st.Query(0, 2)) // 2+3+10 = 15

	st.Update(5, 15)                                                                       // Change arr[5] from 11 to 15
	fmt.Printf("After updating index 5 to 15 - Sum of range [0, 5]: %d\n", st.Query(0, 5)) // 2+3+10+7+9+15 = 46
	fmt.Println()

	// Display final tree structure
	fmt.Println("Final segment tree after updates:")
	st.Display()
}
