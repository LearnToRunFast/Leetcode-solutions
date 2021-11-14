/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (72.30%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    140.9K
 * Total Submissions: 194.9K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 * 回文串 是正着读和反着读都一样的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func dfs(s string, startIndex int, path []string, res *[][]string) {
	if startIndex == len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}
	for i := startIndex; i < len(s); i++ {
		subStr := s[startIndex : i+1]
		if isPalindrome(subStr) {
			dfs(s, i+1, append(path, subStr), res)
		}
	}
}
func partition(s string) [][]string {
	res := make([][]string, 0)
	if len(s) == 0 {
		return res
	}
	dfs(s, 0, []string{}, &res)
	return res
}

// @lc code=end

