/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (51.76%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    69.8K
 * Total Submissions: 134.8K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 *
 *
 * 示例 1：
 *
 * 输入："hello"
 * 输出："holle"
 *
 *
 * 示例 2：
 *
 * 输入："leetcode"
 * 输出："leotcede"
 *
 *
 *
 * 提示：
 *
 *
 * 元音字母不包含字母 "y" 。
 *
 *
 */

// @lc code=start
func isVowels(c rune) bool {
	c = unicode.ToLower(c)
	switch c {
	case 'a':
	case 'e':
	case 'i':
	case 'o':
	case 'u':
	default:
		return false
	}
	return true
}
func reverseVowels(s string) string {
	lo, hi := 0, len(s)-1
	ans := make([]byte, len(s))
	for lo <= hi {
		for lo <= hi && !isVowels(rune(s[lo])) {
			ans[lo] = s[lo]
			lo++
		}
		for lo <= hi && !isVowels(rune(s[hi])) {
			ans[hi] = s[hi]
			hi--
		}
		if lo <= hi {
			ans[lo], ans[hi] = s[hi], s[lo]
			lo++
			hi--
		}
	}
	return string(ans)
}

// @lc code=end

