package main

import "fmt"

// Graph represents a directed graph using an adjacency list
type Graph struct {
	Vertices int
	List     map[int][]int
}

// NewGraph creates a new graph with v vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		List:     make(map[int][]int),
	}
}

// AddEdge adds a directed edge to the graph
func (g *Graph) AddEdge(src, dest int) {
	g.List[src] = append(g.List[src], dest)
}

// TopologicalSort performs a topological sort of the graph
func (g *Graph) TopologicalSort() []int {
	// Kahn's algorithm
	inDegree := make([]int, g.Vertices)
	for _, adjacentVertices := range g.List {
		for _, vertex := range adjacentVertices {
			inDegree[vertex]++
		}
	}

	queue := []int{}
	for i := 0; i < g.Vertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)

		for _, v := range g.List[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != g.Vertices {
		return nil // Graph has a cycle
	}

	return result
}

func main() {
	// Create a graph
	g := NewGraph(6)
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	fmt.Println("Topological Sort:")
	topologicalOrder := g.TopologicalSort()
	if topologicalOrder != nil {
		fmt.Println(topologicalOrder)
	} else {
		fmt.Println("Graph has a cycle!")
	}
}
