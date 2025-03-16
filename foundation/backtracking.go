package main

func generateSubsets(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, start int, path []int, result *[][]int) {
	// add the subset to the result
	*result = append(*result, append([]int{}, path...))

	// start from the current index
	for i := start; i < len(nums); i++ {
		// add the current element to the path
		path = append(path, nums[i])
		// move to the next element
		backtrack(nums, i+1, path, result)
		// backtrack to the previous state
		path = path[:len(path)-1]
	}
}
