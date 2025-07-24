package main

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
