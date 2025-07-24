package main

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
