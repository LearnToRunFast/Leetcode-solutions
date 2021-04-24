/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 *
 * https://leetcode-cn.com/problems/number-of-digit-one/description/
 *
 * algorithms
 * Hard (39.44%)
 * Likes:    211
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 37.3K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 13
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 *
 *
 */

// @lc code=start
func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
func countDigitOne(n int) int {
	var count int64
	var i int64
	N := int64(n)
	for i = 1; i <= N; i *= 10 {
		var div int64 = i * 10
		count += (N / div) * i
		count += min(i, max(0, N%div-i+1))
	}
	return int(count)
}

// @lc code=end

