/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package sollution

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// @lc code=end
