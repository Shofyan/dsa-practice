package main

import (
	"fmt"
)

// DSU represents the Disjoint Set Union data structure
type DSU struct {
	parent []int
	rank   []int
	size   int
}

// NewDSU creates a new DSU with n elements (0 to n-1)
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	// Initialize each element as its own parent
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &DSU{
		parent: parent,
		rank:   rank,
		size:   n,
	}
}

// Find returns the root of the set containing x with path compression
func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x]) // Path compression
	}
	return dsu.parent[x]
}

// Union merges the sets containing x and y using union by rank
func (dsu *DSU) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)

	// Already in the same set
	if rootX == rootY {
		return false
	}

	// Union by rank: attach smaller rank tree under root of higher rank tree
	if dsu.rank[rootX] < dsu.rank[rootY] {
		dsu.parent[rootX] = rootY
	} else if dsu.rank[rootX] > dsu.rank[rootY] {
		dsu.parent[rootY] = rootX
	} else {
		dsu.parent[rootY] = rootX
		dsu.rank[rootX]++
	}

	return true
}

// IsConnected checks if x and y are in the same set
func (dsu *DSU) IsConnected(x, y int) bool {
	return dsu.Find(x) == dsu.Find(y)
}

// GetComponents returns all connected components
func (dsu *DSU) GetComponents() map[int][]int {
	components := make(map[int][]int)

	for i := 0; i < dsu.size; i++ {
		root := dsu.Find(i)
		components[root] = append(components[root], i)
	}

	return components
}

// Display prints the current state of DSU
func (dsu *DSU) Display() {
	fmt.Println("DSU State:")
	fmt.Printf("Parent: %v\n", dsu.parent)
	fmt.Printf("Rank:   %v\n", dsu.rank)

	components := dsu.GetComponents()
	fmt.Println("Connected Components:")
	compNum := 1
	for root, members := range components {
		fmt.Printf("Component %d (root: %d): %v\n", compNum, root, members)
		compNum++
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Disjoint Set Union (DSU) Implementation ===\n")

	// Create DSU with 8 elements (0-7)
	dsu := NewDSU(8)
	fmt.Println("Initial state with 8 elements:")
	dsu.Display()

	// Test union operations
	fmt.Println("Performing union operations:")

	fmt.Println("Union(0, 1):")
	dsu.Union(0, 1)
	dsu.Display()

	fmt.Println("Union(2, 3):")
	dsu.Union(2, 3)
	dsu.Display()

	fmt.Println("Union(4, 5):")
	dsu.Union(4, 5)
	dsu.Display()

	fmt.Println("Union(1, 3): (connecting components {0,1} and {2,3})")
	dsu.Union(1, 3)
	dsu.Display()

	fmt.Println("Union(6, 7):")
	dsu.Union(6, 7)
	dsu.Display()

	// Test connectivity
	fmt.Println("Connectivity Tests:")
	fmt.Printf("IsConnected(0, 3): %t\n", dsu.IsConnected(0, 3)) // true
	fmt.Printf("IsConnected(0, 4): %t\n", dsu.IsConnected(0, 4)) // false
	fmt.Printf("IsConnected(4, 5): %t\n", dsu.IsConnected(4, 5)) // true
	fmt.Printf("IsConnected(6, 7): %t\n", dsu.IsConnected(6, 7)) // true
	fmt.Printf("IsConnected(1, 2): %t\n", dsu.IsConnected(1, 2)) // true
	fmt.Println()

	// More union operations
	fmt.Println("Union(0, 4): (connecting large component with {4,5})")
	dsu.Union(0, 4)
	dsu.Display()

	fmt.Println("Union(5, 6): (connecting everything except isolated elements)")
	dsu.Union(5, 6)
	dsu.Display()

	// Final connectivity tests
	fmt.Println("Final Connectivity Tests:")
	fmt.Printf("IsConnected(0, 7): %t\n", dsu.IsConnected(0, 7)) // true
	fmt.Printf("IsConnected(1, 6): %t\n", dsu.IsConnected(1, 6)) // true
	fmt.Printf("IsConnected(2, 5): %t\n", dsu.IsConnected(2, 5)) // true

	// Test Find operations to show path compression
	fmt.Println("\nFind operations (demonstrating path compression):")
	fmt.Printf("Find(0): %d\n", dsu.Find(0))
	fmt.Printf("Find(3): %d\n", dsu.Find(3))
	fmt.Printf("Find(7): %d\n", dsu.Find(7))

	fmt.Println("\nFinal state after path compression:")
	dsu.Display()
}
