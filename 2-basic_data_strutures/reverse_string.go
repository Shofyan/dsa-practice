package main

import "fmt"

// reverseString reverses a slice of bytes in-place.
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		// Swap the characters at the left and right pointers.
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	str1 := []byte("hello")
	reverseString(str1)
	fmt.Printf("Reversed string: %s
", str1) // Expected output: olleh

	str2 := []byte("Hannah")
	reverseString(str2)
	fmt.Printf("Reversed string: %s
", str2) // Expected output: hannaH
}
