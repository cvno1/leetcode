/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package solution

// @lc code=start
func myAtoi(s string) int {
	sighAllowed, whiteAlloewed, sign, resInts := true, true, 1, []int{}

	for _, c := range s {
		// remove leading white space
		if c == ' ' && whiteAlloewed {
			continue
		}

		// get sigh
		if sighAllowed {
			if c == '+' {
				sign = 1
				sighAllowed = false
				whiteAlloewed = false
				continue
			} else if c == '-' {
				sign = -1
				sighAllowed = false
				whiteAlloewed = false
				continue
			}
		}

		if c < '0' || c > '9' {
			break
		}

		sighAllowed = false
		whiteAlloewed = false

		// remove leading 0
		if len(resInts) == 0 && c == '0' {
			continue
		}

		resInts = append(resInts, int(c-48))
	}

	var max int

	if sign == 1 {
		max = 1<<31 - 1
	} else {
		max = 1 << 31
	}

	res := 0
	if len(resInts) > 10 {
		return sign * max
	} else if len(resInts) == 10 {
		for i := 0; i < len(resInts); i++ {
			if i <= len(resInts)-2 {
				res = res*10 + resInts[i]
			} else {
				if res < max/10 || (res == max/10 && resInts[i] <= max%10) {
					return sign * (res*10 + resInts[i])
				} else {
					return sign * max
				}
			}
		}
	} else {
		for i := 0; i < len(resInts); i++ {
			res = res*10 + resInts[i]
		}
	}

	return sign * res
}

// @lc code=end
