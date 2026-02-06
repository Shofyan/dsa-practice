package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// map sorted string -> list of anagram
	anagramMap := make(map[string][]string)

	// looping di strinng strs
	for _, s := range strs {
		// convert string to character then sort
		chars := []rune(s)
		slices.Sort(chars)

		key := string(chars)
		anagramMap[key] = append(anagramMap[key], s)

	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
