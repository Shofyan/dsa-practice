package main

import "fmt"

// twoSum finds the indices of two numbers that add up to the target.
func twoSum(nums []int, target int) []int {
	// Create a map to store numbers we've seen and their indices.
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// Check if the complement exists in our map.
		if j, ok := numMap[complement]; ok {
			// If it exists, we found our pair.
			return []int{j, i}
		}
		// If the complement is not found, add the current number and its index to the map.
		numMap[num] = i
	}

	// Return an empty slice if no solution is found (though the problem guarantees one).
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Indices of the two numbers are: %v
", result) // Expected output: [0, 1]

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("Indices of the two numbers are: %v
", result2) // Expected output: [1, 2]
}
