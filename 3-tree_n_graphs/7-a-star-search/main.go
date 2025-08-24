package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a node in the graph for A* search
type Node struct {
	ID         int
	X, Y       int // Coordinates for heuristic
	G, H, F    float64
	Parent     *Node
	Edges      []*Edge
	index      int // Index in the priority queue
}

// Edge represents a connection between nodes
type Edge struct {
	To     *Node
	Weight float64
}

// PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an element to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

// Pop removes and returns the element with the highest priority (lowest F score)
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// update modifies the priority and value of a Node in the queue.
func (pq *PriorityQueue) update(node *Node, g, h, f float64) {
	node.G = g
	node.H = h
	node.F = f
	heap.Fix(pq, node.index)
}

// heuristic calculates the Euclidean distance between two nodes
func heuristic(a, b *Node) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

// AStar finds the shortest path from start to goal using A* algorithm
func AStar(start, goal *Node) ([]*Node, float64) {
	openSet := make(PriorityQueue, 0)
	heap.Push(&openSet, start)

	start.G = 0
	start.H = heuristic(start, goal)
	start.F = start.G + start.H

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current == goal {
			path := []*Node{}
			for temp := current; temp != nil; temp = temp.Parent {
				path = append([]*Node{temp}, path...)
			}
			return path, current.F
		}

		for _, edge := range current.Edges {
			neighbor := edge.To
			tentativeG := current.G + edge.Weight

			if tentativeG < neighbor.G {
				neighbor.Parent = current
				neighbor.G = tentativeG
				neighbor.H = heuristic(neighbor, goal)
				neighbor.F = neighbor.G + neighbor.H
				// Check if neighbor is in openSet
				inOpenSet := false
				for _, n := range openSet {
					if n == neighbor {
						inOpenSet = true
						break
					}
				}
				if !inOpenSet {
					heap.Push(&openSet, neighbor)
				} else {
					openSet.update(neighbor, neighbor.G, neighbor.H, neighbor.F)
				}
			}
		}
	}

	return nil, 0 // No path found
}

func main() {
	// Create nodes
	nodes := []*Node{
		{ID: 0, X: 0, Y: 0, G: math.Inf(1), H: math.Inf(1), F: math.Inf(1)},
		{ID: 1, X: 1, Y: 2, G: math.Inf(1), H: math.Inf(1), F: math.Inf(1)},
		{ID: 2, X: 2, Y: 1, G: math.Inf(1), H: math.Inf(1), F: math.Inf(1)},
		{ID: 3, X: 3, Y: 3, G: math.Inf(1), H: math.Inf(1), F: math.Inf(1)},
		{ID: 4, X: 4, Y: 0, G: math.Inf(1), H: math.Inf(1), F: math.Inf(1)},
	}

	// Create edges
	nodes[0].Edges = []*Edge{{To: nodes[1], Weight: 2.2}, {To: nodes[2], Weight: 2.2}}
	nodes[1].Edges = []*Edge{{To: nodes[0], Weight: 2.2}, {To: nodes[3], Weight: 2.8}}
	nodes[2].Edges = []*Edge{{To: nodes[0], Weight: 2.2}, {To: nodes[3], Weight: 2.2}, {To: nodes[4], Weight: 4.5}}
	nodes[3].Edges = []*Edge{{To: nodes[1], Weight: 2.8}, {To: nodes[2], Weight: 2.2}, {To: nodes[4], Weight: 2.2}}
	nodes[4].Edges = []*Edge{{To: nodes[2], Weight: 4.5}, {To: nodes[3], Weight: 2.2}}

	startNode := nodes[0]
	goalNode := nodes[4]

	path, distance := AStar(startNode, goalNode)

	if path != nil {
		fmt.Printf("Shortest path from %d to %d:\n", startNode.ID, goalNode.ID)
		for _, node := range path {
			fmt.Printf(" -> %d", node.ID)
		}
		fmt.Printf("\nTotal distance: %.2f\n", distance)
	} else {
		fmt.Printf("No path found from %d to %d\n", startNode.ID, goalNode.ID)
	}
}
