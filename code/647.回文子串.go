/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (65.74%)
 * Likes:    712
 * Dislikes: 0
 * Total Accepted:    130.3K
 * Total Submissions: 198.2K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
func countSub(s string, l, r, n int) int {
	count := 0
	for l >= 0 && r < n && s[l] == s[r] {
		l--
		r++
		count++
	}
	return count
}
func countSubstrings(s string) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		count += countSub(s, i, i, n)   // i as center
		count += countSub(s, i, i+1, n) // i,i+1 as center
	}
	return count
}

// @lc code=end

