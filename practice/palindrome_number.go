package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	p0 := 0
	p1 := len(s) - 1

	for p0 < p1 {
		if s[p0] != s[p1] {
			return false
		}
		p0++
		p1--
	}
	return true
}
