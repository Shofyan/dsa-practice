package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a connection between two nodes with a weight (cost).
type Edge struct {
	target, weight int
}

// Graph is a map where each node points to a list of edges (connections to other nodes).
type Graph map[int][]Edge

// Item represents a node with its cost used in the priority queue.
type Item struct {
	node, cost int
}

// PriorityQueue implements a min-heap based on the cost of visiting a node.
type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest path from a starting node to all other nodes in the graph.
func Dijkstra(graph Graph, start int) map[int]int {
	// Initialize distances to all nodes as infinite.
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt64
	}
	// Set the starting node distance to 0.
	dist[start] = 0

	// Create a priority queue and push the start node.
	pq := &PriorityQueue{{node: start, cost: 0}}
	heap.Init(pq)

	// Process the queue until all reachable nodes are visited.
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		// Skip if a shorter distance has already been found.
		if current.cost > dist[current.node] {
			continue
		}
		// Explore all neighboring nodes.
		for _, edge := range graph[current.node] {
			newDist := dist[current.node] + edge.weight
			// If a shorter path is found, update and push to the queue.
			if newDist < dist[edge.target] {
				dist[edge.target] = newDist
				heap.Push(pq, Item{node: edge.target, cost: newDist})
			}
		}
	}
	return dist
}

// PrintGraph displays the graph structure in a readable format.
func PrintGraph(graph Graph) {
	fmt.Println("Graph representation:")
	for node, edges := range graph {
		fmt.Printf("Node %d: ", node)
		for _, edge := range edges {
			fmt.Printf("(%d -> %d, weight: %d) ", node, edge.target, edge.weight)
		}
		fmt.Println()
	}
}
