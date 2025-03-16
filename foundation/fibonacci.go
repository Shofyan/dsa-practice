package main

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// fibonacci memoization∫
func fibonacciMemoization(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacciMemoization(n-1, memo) + fibonacciMemoization(n-2, memo)
	return memo[n]
}

// fibonacci bottom up DP
func fibonacciBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
