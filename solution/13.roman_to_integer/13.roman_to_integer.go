/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package sollution

// @lc code=start
func romanToInt(s string) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res, i := 0, 0
	s += " "
	for len(s) > 1 {
		for s[:len(symbols[i])] != symbols[i] {
			i++
		}
		res += values[i]
		s = s[len(symbols[i]):]
	}

	return res

}

// @lc code=end
