/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package sollution

// @lc code=start
func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		max = longest(max, maxPalindrome(s, i, i))
		max = longest(max, maxPalindrome(s, i, i+1))
	}
	return max
}

func maxPalindrome(s string, i, j int) string {
	res := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res = s[i : j+1]
		i--
		j++
	}
	return res
}

func longest(s1, s2 string) string {
	if len(s1) >= len(s2) {
		return s1
	} else {
		return s2
	}
}

// @lc code=end
