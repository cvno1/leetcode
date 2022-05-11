/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package solution

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	total := len(nums1) + len(nums2)
	l := make([]int, total)
	for k := 0; k <= total; k++ {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				l[k] = nums1[i]
				i++
			} else {
				l[k] = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			l[k] = nums1[i]
			i++
		} else if j < len(nums2) {
			l[k] = nums2[j]
			j++
		}
	}

	if total%2 == 0 {
		return float64(l[total>>1]+l[total>>1-1]) / 2
	} else {
		return float64(l[total>>1])
	}
}

// @lc code=end
