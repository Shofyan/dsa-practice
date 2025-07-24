package main

import "fmt"

// MyStack is a LIFO stack implemented with two queues.
type MyStack struct {
	q1 []int
	q2 []int
}

// Constructor initializes the stack.
func Constructor() MyStack {
	return MyStack{}
}

// Push element x onto stack.
func (s *MyStack) Push(x int) {
	// Add the new element to the empty queue.
	s.q2 = append(s.q2, x)
	// Move all elements from q1 to q2.
	for len(s.q1) > 0 {
		s.q2 = append(s.q2, s.q1[0])
		s.q1 = s.q1[1:]
	}
	// Swap the queues.
	s.q1, s.q2 = s.q2, s.q1
}

// Pop removes the element on top of the stack and returns it.
func (s *MyStack) Pop() int {
	top := s.q1[0]
	s.q1 = s.q1[1:]
	return top
}

// Top gets the top element.
func (s *MyStack) Top() int {
	return s.q1[0]
}

// Empty returns whether the stack is empty.
func (s *MyStack) Empty() bool {
	return len(s.q1) == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Printf("Top element is: %d
", stack.Top())   // returns 2
	fmt.Printf("Popped element is: %d
", stack.Pop()) // returns 2
	fmt.Printf("Is stack empty? %v
", stack.Empty()) // returns false
}
