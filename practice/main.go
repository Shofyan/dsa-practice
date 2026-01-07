package main

import "fmt"

func main() {
  
     	println("valid parentheses Test Cases:")
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"", true},          // edge case (optional)
		{"(((", false},
		{"{[()]}", true},
	}

	for i, test := range tests {
		result := isValid(test.input)
		status := "FAIL"
		if result == test.expected {
			status = "PASS"
		}

		fmt.Printf(
			"Test #%d | input: %-8s | expected: %-5v | got: %-5v | %s\n",
			i+1,
			test.input,
			test.expected,
			result,
			status,
		)
	}

	// lowest common ancestor
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	p := root.Left
	q := root.Right

	lca := lowestCommonAncestor(root, p, q)
	lca = lowestCommonAncestorNorecursion(root, p, q)
	fmt.Printf("Lowest Common Ancestor of %d and %d is: %d\n", p.Val, q.Val, lca.Val)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 7

	findPairs([]int{1, 2, 3, 4, 5}, 5)

	pairs := cariTargetAngka(arr, target)
	fmt.Println(pairs)

	pairsP := findPairsP(arr, target)
	fmt.Println(pairsP)

	// Test Case 1
	println("Task Scheduler Test Cases:")
	tasks1 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n1 := 2
	fmt.Println("Test 1 Output:", leastInterval(tasks1, n1)) // Expected: 8

	// Test Case 2
	tasks2 := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n2 := 1
	fmt.Println("Test 2 Output:", leastInterval(tasks2, n2)) // Expected: 6

	// Test Case 3
	tasks3 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n3 := 3
	fmt.Println("Test 3 Output:", leastInterval(tasks3, n3)) // Expected: 10

	// Edge Case: No cooldown
	tasks4 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n4 := 0
	fmt.Println("Test 4 Output:", leastInterval(tasks4, n4)) // Expected: 6

	// Edge Case: Single task
	tasks5 := []byte{'A'}
	n5 := 5
	fmt.Println("Test 5 Output:", leastInterval(tasks5, n5)) // Expected: 1

}
