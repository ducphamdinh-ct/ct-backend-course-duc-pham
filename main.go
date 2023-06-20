package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid(")"))      // false
	fmt.Println(isValid("([)"))
}
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			previous := stack[len(stack)-1]
			if previous == '(' {
				stack = append(stack, char)
			} else {
				return false
			}
		case '}':
			previous := stack[len(stack)-1]
			if previous == '{' {
				stack = append(stack, char)
			} else {
				return false
			}
		case ']':
			previous := stack[len(stack)-1]
			if previous == '[' {
				stack = append(stack, char)
			} else {
				return false
			}
		}
	}
	return len(stack)%2 == 0
}
