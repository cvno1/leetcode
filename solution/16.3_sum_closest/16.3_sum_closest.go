/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package sollution

import (
	"math"
	"sort"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i++ {
		left, right := 0, len(nums)-1
		for left < i && right > i {
			total := nums[left] + nums[i] + nums[right]
			if math.Abs(float64(total)-float64(target)) < math.Abs(float64(result)-float64(target)) {
				result = total
			}

			if total > target {
				right--
			} else if total == target {
				return total
			} else {
				left++
			}
		}
	}
	return result
}

// @lc code=end
