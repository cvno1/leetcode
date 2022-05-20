/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 */
package sollution

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 1; i < len(nums)-1; i++ {
		left, right := 0, len(nums)-1

		if i > 1 && nums[i] == nums[i-1] {
			left = i - 1
		}

		for left < i && right > i {

			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(result, []int{nums[left], nums[i], nums[right]})
				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// @lc code=end
