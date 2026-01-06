package main

func leastInterval(tasks []byte, n int) int {

	// Step 1: Count frequency of each task (A-Z)
	// Since tasks are bytes representing 'A' to 'Z',
	// we can map them directly into a fixed-size array.
	freq := make([]int, 26)

	for _, task := range tasks {
		freq[task-'A']++
	}

	// Step 2: Find the maximum frequency among all tasks
	// This task will determine the skeleton of the schedule.
	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// Step 3: Count how many tasks appear maxFreq times
	// Example: A=3, B=3 → maxCount = 2
	maxCount := 0
	for _, f := range freq {
		if f == maxFreq {
			maxCount++
		}
	}

	/*
	   Step 4: Compute the minimum number of CPU intervals

	   Idea:
	   - The most frequent task creates (maxFreq - 1) gaps
	   - Each gap needs n cooldown intervals
	   - Each block has (n + 1) slots
	   - Add maxCount for the last block

	   Formula:
	   (maxFreq - 1) * (n + 1) + maxCount
	*/
	partCount := maxFreq - 1
	partLength := n + 1
	minIntervals := partCount*partLength + maxCount

	// Step 5: The result must be at least the total number of tasks
	// If we have enough tasks to fill the idle slots,
	// then no idle time is required.
	if minIntervals < len(tasks) {
		return len(tasks)
	}

	return minIntervals
}
