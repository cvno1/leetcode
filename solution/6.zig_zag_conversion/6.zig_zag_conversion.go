/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
package solution

import "strings"

// @lc code=start
func convert(s string, n int) string {
	if n == 1 {
		return s
	}
	res := make([]string, n)
	grouplen := 2*n - 2

	for i, char := range s {
		x := i % grouplen
		res[min(x, grouplen-x)] += string(char)
	}
	return strings.Join(res, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
