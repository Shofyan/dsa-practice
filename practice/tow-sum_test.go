package main

import (
	"fmt"
	"testing"
)

func FindTargetPairs(arr []int, target int) [][]int {
	result := [][]int{}
	seen := make(map[int]bool)
	pairs := make(map[string]bool)

	for _, angka := range arr {

		need := target - angka

		if seen[need] {
			a, b := need, angka
			if a > b {
				a, b = b, a
			}

			key := fmt.Sprintf("%d-%d", a, b)

			if !pairs[key] {
				pairs[key] = true
				result = append(result, []int{a, b})
			}

		}
		seen[angka] = true

	}

	return result
}

func TestFindTargetPairs(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	pairs := FindTargetPairs(arr, 7)

	t.Log(pairs)
	fmt.Println(pairs)

}
