/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package solution

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0, 32)

	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}

	for i := 0; i <= (len(nums)-1)>>1; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}

	return true
}

// @lc code=end
