// Package main adalah program utama untuk demonstrasi representasi graf menggunakan adjacency matrix dan adjacency list.
package main

import "fmt"

// Graph menggunakan Adjacency Matrix (Matriks Ketetanggaan)
type AdjacencyMatrixGraph struct {
	// Jumlah simpul (vertex) dalam graf
	Vertices int
	// Matriks ketetanggaan untuk menyimpan hubungan antar simpul
	Matrix [][]int
}

// NewAdjacencyMatrixGraph membuat graf baru dengan jumlah simpul tertentu
func NewAdjacencyMatrixGraph(vertices int) *AdjacencyMatrixGraph {
	// Membuat matriks 2 dimensi berukuran vertices x vertices
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &AdjacencyMatrixGraph{
		Vertices: vertices,
		Matrix:   matrix,
	}
}

// AddEdge menambahkan sisi (edge) antara dua simpul pada graf
func (g *AdjacencyMatrixGraph) AddEdge(src, dest int) {
	g.Matrix[src][dest] = 1 // Tandai ada sisi dari src ke dest
	g.Matrix[dest][src] = 1 // Untuk graf tak berarah, tandai juga dari dest ke src
}

// Print menampilkan matriks ketetanggaan ke layar
func (g *AdjacencyMatrixGraph) Print() {
	fmt.Println("Adjacency Matrix:")
	for _, row := range g.Matrix {
		fmt.Println(row)
	}
}

// Graph menggunakan Adjacency List (List Ketetanggaan)
type AdjacencyListGraph struct {
	// Jumlah simpul (vertex) dalam graf
	Vertices int
	// Map untuk menyimpan list ketetanggaan tiap simpul
	List map[int][]int
}

// NewAdjacencyListGraph membuat graf baru dengan jumlah simpul tertentu
func NewAdjacencyListGraph(vertices int) *AdjacencyListGraph {
	// Membuat map kosong untuk adjacency list
	return &AdjacencyListGraph{
		Vertices: vertices,
		List:     make(map[int][]int),
	}
}

// AddEdge menambahkan sisi (edge) antara dua simpul pada graf
func (g *AdjacencyListGraph) AddEdge(src, dest int) {
	g.List[src] = append(g.List[src], dest)  // Tambahkan dest ke list src
	g.List[dest] = append(g.List[dest], src) // Untuk graf tak berarah, tambahkan src ke list dest
}

// Print menampilkan adjacency list ke layar
func (g *AdjacencyListGraph) Print() {
	fmt.Println("Adjacency List:")
	for vertex, edges := range g.List {
		fmt.Printf("Vertex %d: %v\n", vertex, edges)
	}
}

func main() {
	// Contoh penggunaan graf dengan adjacency matrix
	v := 5 // Jumlah simpul
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

	// Contoh penggunaan graf dengan adjacency list
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
