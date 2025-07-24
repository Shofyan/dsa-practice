package main

func nQueen(n int) [][]int {
	result := [][]int{}
	backtrack_queen(n, 0, []int{}, &result)
	return result
}

func backtrack_queen(n int, row int, path []int, result *[][]int) {
	if row == n {
		*result = append(*result, append([]int{}, path...))
		return
	}
	for col := 0; col < n; col++ {
		if isValid_queen(row, col, path) {
			path = append(path, col)
			backtrack_queen(n, row+1, path, result)
			path = path[:len(path)-1]
		}
	}
}

func isValid_queen(row int, col int, path []int) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == col || abs(row-i) == abs(col-path[i]) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
