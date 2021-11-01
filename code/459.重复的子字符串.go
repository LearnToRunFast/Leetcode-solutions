/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (51.19%)
 * Likes:    512
 * Dislikes: 0
 * Total Accepted:    73.5K
 * Total Submissions: 143.5K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */

// @lc code=start
// func repeatedSubstringPattern(s string) bool {
// 	newS := s + s
// 	return strings.Contains(newS[1:len(newS)-1], s)
// }
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	prefix := make([]int, n)
	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = prefix[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		prefix[i] = j
	}
	subLen := n - prefix[n-1]
	if prefix[n-1] != 0 && n%subLen == 0 {
		return true
	}
	return false
}

// @lc code=end

