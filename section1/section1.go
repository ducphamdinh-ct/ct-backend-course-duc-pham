package section1

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
