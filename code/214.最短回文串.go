/*
 * @lc app=leetcode.cn id=214 lang=golang
 *
 * [214] 最短回文串
 *
 * https://leetcode-cn.com/problems/shortest-palindrome/description/
 *
 * algorithms
 * Hard (36.77%)
 * Likes:    329
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 73.6K
 * Testcase Example:  '"aacecaaa"'
 *
 * 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aacecaaa"
 * 输出："aaacecaaa"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd"
 * 输出："dcbabcd"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 仅由小写英文字母组成
 *
 *
 */
package main

import (
	"strings"
)

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverse(s string) string {
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
func isPalindrome(i, j int, dp []int) bool {
	m := len(dp)
	l := 2*i + 1
	r := 2*j + 1
	mid := l + (r-l)/2
	if mid >= 0 && mid < m && r-l+1 <= dp[mid] {
		return true
	}
	return false
}
func shortestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	m := 2*n + 1
	newS := make([]byte, m)
	for i := range newS {
		if i&1 == 1 {
			newS[i] = s[i>>1]
		} else {
			newS[i] = '$'
		}
	}

	l, r := 0, -1
	dp := make([]int, m)
	for i := range newS {
		// in range
		if i <= r {
			dp[i] = min((r-i)*2+1, dp[l+r-i])
		} else {
			dp[i] = 1
		}

		temp_l := i - dp[i]/2 - 1
		temp_r := i + dp[i]/2 + 1

		for temp_l >= 0 && temp_r < m && newS[temp_l] == newS[temp_r] {
			temp_l--
			temp_r++
		}
		dp[i] = temp_r - temp_l - 1
		if temp_r > r {
			l = temp_l + 1
			r = temp_r - 1
		}
	}
	for i := m - 1; i >= 0; i-- {
		if isPalindrome(0, i, dp) {
			reversed := reverse(s[i+1:])
			return reversed + s
		}
	}
	return ""
}

// @lc code=end
