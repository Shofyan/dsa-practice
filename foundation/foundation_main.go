package main

import (
	"fmt"
)

func main() {
	println("Factorial result: ", Factorial(4))

	println("FactorialIterative result: ", FactorialIterative(4))

	println("Fibonacci result: ", fibonacci(4))

	println("FibonacciMemoization result: ", fibonacciMemoization(4, make(map[int]int)))

	println("FibonacciBottomUp result: ", fibonacciBottomUp(4))

	set := []int{1, 2, 3}
	subsets := generateSubsets(set)
	fmt.Println("All subsets:", subsets)

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	printBoard(board)

	println("N-Queen result: ")
	//loop for pirnt the result

	for _, v := range nQueen(5) {
		fmt.Println(v)
	}

	// Define a graph with weighted edges.
	graph := Graph{
		0: {{1, 10}, {2, 15}, {3, 20}},
		1: {{0, 10}, {3, 25}, {4, 30}},
		2: {{0, 15}, {4, 10}},
		3: {{0, 20}, {1, 25}, {4, 20}},
		4: {{1, 30}, {2, 10}, {3, 20}},
	}

	// Print the graph before running Dijkstra's algorithm.
	PrintGraph(graph)

	// Run Dijkstra's algorithm from a starting city (node 0).
	startCity := 0
	distances := Dijkstra(graph, startCity)
	fmt.Printf("\nShortest distances from city %d:\n", startCity)
	for city, dist := range distances {
		fmt.Printf("City %d: %d\n", city, dist)
	}

}
