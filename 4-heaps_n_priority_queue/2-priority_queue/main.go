package main

import (
	"fmt"
)

type PriorityQueue struct {
	heap []int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap: []int{}}
}

func (pq *PriorityQueue) Push(x int) {
	pq.heap = append(pq.heap, x)
	pq.up(len(pq.heap) - 1)
}

func (pq *PriorityQueue) Pop() (int, bool) {
	if len(pq.heap) == 0 {
		return 0, false
	}
	min := pq.heap[0]
	last := pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	if len(pq.heap) > 0 {
		pq.heap[0] = last
		pq.down(0)
	}
	return min, true
}

func (pq *PriorityQueue) Peek() (int, bool) {
	if len(pq.heap) == 0 {
		return 0, false
	}
	return pq.heap[0], true
}

func (pq *PriorityQueue) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.heap[i] >= pq.heap[parent] {
			break
		}
		pq.heap[i], pq.heap[parent] = pq.heap[parent], pq.heap[i]
		i = parent
	}
}

func (pq *PriorityQueue) down(i int) {
	n := len(pq.heap)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < n && pq.heap[left] < pq.heap[smallest] {
			smallest = left
		}
		if right < n && pq.heap[right] < pq.heap[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		pq.heap[i], pq.heap[smallest] = pq.heap[smallest], pq.heap[i]
		i = smallest
	}
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap) == 0
}

func main() {
	pq := NewPriorityQueue()
	pq.Push(5)
	pq.Push(3)
	pq.Push(8)
	pq.Push(1)
	fmt.Println("Pop order:")
	for !pq.IsEmpty() {
		v, _ := pq.Pop()
		fmt.Println(v)
	}
}
