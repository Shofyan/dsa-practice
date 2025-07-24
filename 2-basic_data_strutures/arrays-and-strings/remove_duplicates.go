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
	fmt.Printf("New length: %d, Array: %v", length1, nums1[:length1]) // Expected output: 2, [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length2 := removeDuplicates(nums2)
	fmt.Printf("New length: %d, Array: %v", length2, nums2[:length2]) // Expected output: 5, [0 1 2 3 4]

	str1 := []byte("hello")
	reverseString(str1)
	fmt.Printf("Reversed string: %s", str1) // Expected output: olleh

	str2 := []byte("Hannah")
	reverseString(str2)
	fmt.Printf("Reversed string: %s", str2) // Expected output: hannaH

	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Indices of the two numbers are: %v", result) // Expected output: [0, 1]

	nums2 = []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("Indices of the two numbers are: %v", result2) // Expected output: [1, 2]
}
