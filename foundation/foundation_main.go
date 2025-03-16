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

}
