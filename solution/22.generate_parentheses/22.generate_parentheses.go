/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package sollution

// @lc code=start
var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	if n == 0 {
		return result
	}

	findGenerateParenthesis(n, n, "")

	return result

}

func findGenerateParenthesis(lcount, rcount int, str string) {
	if lcount == 0 && rcount == 0 {
		result = append(result, str)
		return
	}
	if lcount > 0 {
		findGenerateParenthesis(lcount-1, rcount, str+"(")
	}
	if rcount > 0 && rcount > lcount {
		findGenerateParenthesis(lcount, rcount-1, str+")")
	}
}

// @lc code=end
