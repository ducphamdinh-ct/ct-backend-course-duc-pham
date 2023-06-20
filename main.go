package main

import "fmt"

func main() {
	fmt.Println(isValidParentheses("()"))       // true
	fmt.Println(isValidParentheses("()[]{}"))   // true
	fmt.Println(isValidParentheses("(]"))       // false
	fmt.Println(isValidParentheses("([)]"))     // false
	fmt.Println(isValidParentheses("{[]}"))     // true
	fmt.Println(isValidParentheses("(([])[])")) // true

	fmt.Println(containsDuplicate([]int{1, 2, 3}))          // false
	fmt.Println(containsDuplicate([]int{4, 5, 6}))          // false
	fmt.Println(containsDuplicate([]int{1, 1, 3, 4, 5, 4})) // true
	fmt.Println(containsDuplicate([]int{2, 1, 3, 4, 5}))    // false
	fmt.Println(containsDuplicate([]int{9, 8, 6, 9}))       // true
}

func containsDuplicate(nums []int) bool {
	storage := map[int]bool{}
	for _, num := range nums {
		if storage[num] {
			return true
		}
		storage[num] = true
	}
	return false
}
func isValidParentheses(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			previous := stack[len(stack)-1]
			if previous == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			previous := stack[len(stack)-1]
			if previous == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			previous := stack[len(stack)-1]
			if previous == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
