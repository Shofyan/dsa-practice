package main

import "fmt"

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

}
