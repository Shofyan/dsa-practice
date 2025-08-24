package main

import "fmt"

// AdjacencyListGraph represents a graph using an adjacency list
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
	g.List[dest] = append(g.List[dest], src) // For an undirected graph
}

// BFS performs Breadth-First Search starting from a given vertex
func (g *AdjacencyListGraph) BFS(startVertex int) {
	visited := make([]bool, g.Vertices)
	queue := []int{startVertex}
	visited[startVertex] = true

	fmt.Print("BFS Traversal: ")
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex)

		for _, adjacentVertex := range g.List[vertex] {
			if !visited[adjacentVertex] {
				visited[adjacentVertex] = true
				queue = append(queue, adjacentVertex)
			}
		}
	}
	fmt.Println()
}

// DFS performs Depth-First Search starting from a given vertex
func (g *AdjacencyListGraph) DFS(startVertex int) {
	visited := make([]bool, g.Vertices)
	fmt.Print("DFS Traversal: ")
	g.dfsUtil(startVertex, visited)
	fmt.Println()
}

// dfsUtil is a recursive helper function for DFS
func (g *AdjacencyListGraph) dfsUtil(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	for _, adjacentVertex := range g.List[vertex] {
		if !visited[adjacentVertex] {
			g.dfsUtil(adjacentVertex, visited)
		}
	}
}

func main() {
	v := 5
	graph := NewAdjacencyListGraph(v)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)

	fmt.Println("Graph:")
	for vertex, edges := range graph.List {
		fmt.Printf("Vertex %d: %v\n", vertex, edges)
	}

	fmt.Println()
	graph.BFS(0)
	graph.DFS(0)
}
