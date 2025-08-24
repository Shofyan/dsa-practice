package main

import (
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents a weighted graph using an adjacency list
type Graph struct {
	Vertices int
	List     map[int][]Edge
}

// NewGraph creates a new graph with v vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		List:     make(map[int][]Edge),
	}
}

// AddEdge adds a directed edge with a weight to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.List[from] = append(g.List[from], Edge{To: to, Weight: weight})
}

// Dijkstra finds the shortest paths from a source vertex to all other vertices
func (g *Graph) Dijkstra(src int) []int {
	dist := make([]int, g.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	visited := make([]bool, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		u := -1
		for j := 0; j < g.Vertices; j++ {
			if !visited[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}

		if u == -1 || dist[u] == math.MaxInt32 {
			break
		}

		visited[u] = true

		for _, edge := range g.List[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
			}
		}
	}

	return dist
}

func main() {
	// Create a graph
	g := NewGraph(9)
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 8, 2)
	g.AddEdge(2, 5, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(3, 5, 14)
	g.AddEdge(4, 5, 10)
	g.AddEdge(5, 6, 2)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 6)
	g.AddEdge(7, 8, 7)

	source := 0
	distances := g.Dijkstra(source)

	fmt.Printf("Shortest paths from source vertex %d:\n", source)
	for i, d := range distances {
		fmt.Printf("Vertex %d: Distance %d\n", i, d)
	}
}
