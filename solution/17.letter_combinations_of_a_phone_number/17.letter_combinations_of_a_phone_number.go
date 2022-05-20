/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package sollution

// @lc code=start
var result []string

var dic = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	lettersfunc([]byte{}, []byte(digits))
	return result
}

func lettersfunc(res, digits []byte) {
	if len(digits) == 0 {
		result = append(result, string(res))
		return
	}

	key := digits[0]
	digits = digits[1:]

	for i := 0; i < len(dic[key]); i++ {
		res = append([]byte(res), dic[key][i])
		lettersfunc(res, digits)
		res = res[0 : len(res)-1]
	}
}

// @lc code=end
