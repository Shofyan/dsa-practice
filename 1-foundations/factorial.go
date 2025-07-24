package main

func Factorial(n int) int {

	println("Factorial")
	println("n: ", n)

	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// factorial iterative
func FactorialIterative(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
