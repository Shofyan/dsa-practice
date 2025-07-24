package main

import "fmt"

// removeDuplicates removes duplicates from a sorted slice in-place.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}
	return insertIndex
}

func main() {
	nums1 := []int{1, 1, 2}
	length1 := removeDuplicates(nums1)
	fmt.Printf("New length: %d, Array: %v
", length1, nums1[:length1]) // Expected output: 2, [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length2 := removeDuplicates(nums2)
	fmt.Printf("New length: %d, Array: %v
", length2, nums2[:length2]) // Expected output: 5, [0 1 2 3 4]
}
