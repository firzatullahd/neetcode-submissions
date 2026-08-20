func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					return false
				}
				stack = stack[:len(stack)-1]
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					return false
				}
				stack = stack[:len(stack)-1]
			case '}':
				if len(stack) == 0 || stack[len(stack)-1] != '{' {
					return false
				}
				stack = stack[:len(stack)-1]
			case '(', '[', '{':
				stack = append(stack, v)
		}
	}

	return len(stack) == 0
}