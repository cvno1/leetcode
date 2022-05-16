/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package sollution

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	res := ""

	for i := 0; i < len(strs); i++ {
		if i == 0 {
			res = strs[i]
			continue
		}

		j := 0
		for j < min(len(res), len(strs[i])) && res[j] == strs[i][j] {
			j++
		}
		res = res[:j]
		if len(res) == 0 {
			break
		}
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// @lc code=end
