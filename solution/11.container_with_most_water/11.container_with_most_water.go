/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package sollution

// @lc code=start
func maxArea(h []int) int {
	left, right, max := 0, len(h)-1, 0

	for left < right {
		width := right - left
		lower := 0
		if h[left] > h[right] {
			lower = h[right]
			right--
		} else {
			lower = h[left]
			left++
		}

		if lower*width > max {
			max = lower * width
		}

	}

	return max
}

// @lc code=end
