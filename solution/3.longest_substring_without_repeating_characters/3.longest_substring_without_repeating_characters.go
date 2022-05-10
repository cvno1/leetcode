/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package solution

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	right, left, res := 0, 0, 0
	hashmap := make(map[byte]int, len(s))
	for right < len(s) {
		if idx, ok := hashmap[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		hashmap[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
