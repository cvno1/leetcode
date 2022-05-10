/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] two_sum
 */

package solution

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := hashmap[target-nums[i]]; ok {
			return []int{idx, i}
		}
		hashmap[nums[i]] = i
	}
	return []int{}
}

// @lc code=end
