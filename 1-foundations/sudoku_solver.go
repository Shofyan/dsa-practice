package main

import (
	"fmt"
	"math"
)

func solveSudoku(board [][]byte) {
	solve(board)
}
func solve(board [][]byte) bool {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= byte('0')+byte(n); c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	n := len(board)
	subGridSize := int(math.Sqrt(float64(n)))
	for i := 0; i < n; i++ {
		if board[i][col] == c {
			return false
		}
		if board[row][i] == c {
			return false
		}
		if board[subGridSize*(row/subGridSize)+i/subGridSize][subGridSize*(col/subGridSize)+i%subGridSize] == c {
			return false
		}
	}
	return true
}

func printBoard(board [][]byte) {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}
