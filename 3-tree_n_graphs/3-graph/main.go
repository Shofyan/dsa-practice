package main

import "fmt"

// Graph using Adjacency Matrix
type AdjacencyMatrixGraph struct {
	Vertices int
	Matrix   [][]int
}

// NewAdjacencyMatrixGraph creates a new graph with v vertices
func NewAdjacencyMatrixGraph(vertices int) *AdjacencyMatrixGraph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &AdjacencyMatrixGraph{
		Vertices: vertices,
		Matrix:   matrix,
	}
}

// AddEdge adds an edge to the graph
func (g *AdjacencyMatrixGraph) AddEdge(src, dest int) {
	g.Matrix[src][dest] = 1
	g.Matrix[dest][src] = 1 // For undirected graph
}

// Print prints the adjacency matrix
func (g *AdjacencyMatrixGraph) Print() {
	fmt.Println("Adjacency Matrix:")
	for _, row := range g.Matrix {
		fmt.Println(row)
	}
}

// Graph using Adjacency List
type AdjacencyListGraph struct {
	Vertices int
	List     map[int][]int
}

// NewAdjacencyListGraph creates a new graph with v vertices
func NewAdjacencyListGraph(vertices int) *AdjacencyListGraph {
	return &AdjacencyListGraph{
		Vertices: vertices,
		List:     make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph
func (g *AdjacencyListGraph) AddEdge(src, dest int) {
	g.List[src] = append(g.List[src], dest)
	g.List[dest] = append(g.List[dest], src) // For undirected graph
}

// Print prints the adjacency list
func (g *AdjacencyListGraph) Print() {
	fmt.Println("Adjacency List:")
	for vertex, edges := range g.List {
		fmt.Printf("Vertex %d: %v\n", vertex, edges)
	}
}

func main() {
	// Adjacency Matrix Test
	v := 5
	adjMatrixGraph := NewAdjacencyMatrixGraph(v)
	adjMatrixGraph.AddEdge(0, 1)
	adjMatrixGraph.AddEdge(0, 4)
	adjMatrixGraph.AddEdge(1, 2)
	adjMatrixGraph.AddEdge(1, 3)
	adjMatrixGraph.AddEdge(1, 4)
	adjMatrixGraph.AddEdge(2, 3)
	adjMatrixGraph.AddEdge(3, 4)
	adjMatrixGraph.Print()

	fmt.Println()

	// Adjacency List Test
	adjListGraph := NewAdjacencyListGraph(v)
	adjListGraph.AddEdge(0, 1)
	adjListGraph.AddEdge(0, 4)
	adjListGraph.AddEdge(1, 2)
	adjListGraph.AddEdge(1, 3)
	adjListGraph.AddEdge(1, 4)
	adjListGraph.AddEdge(2, 3)
	adjListGraph.AddEdge(3, 4)
	adjListGraph.Print()
}