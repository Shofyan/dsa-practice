package main

func isValid(s string) bool {
    stack := []rune{}
    
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, c := range s {
        // Opening bracket
        if c == '(' || c == '{' || c == '[' {
            stack = append(stack, c)
            continue
        }

        // Closing bracket
        if len(stack) == 0 {
            return false
        }

        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if pairs[c] != top {
            return false
        }
    }

    return len(stack) == 0
}
