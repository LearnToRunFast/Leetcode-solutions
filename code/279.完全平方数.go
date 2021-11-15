/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (63.35%)
 * Likes:    1137
 * Dislikes: 0
 * Total Accepted:    213.5K
 * Total Submissions: 336.8K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 *
 * 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
 *
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：3
 * 解释：12 = 4 + 4 + 4
 *
 * 示例 2：
 *
 *
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func numSquares(n int) int {
	if n < 3 {
		return n
	}
	// define dp[i] as the minimum number of perfect squares to sum to i
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// possible max value is i.
		dp[i] = i
	}
	for i := 1; i*i <= n; i++ {
		value := i * i
		for j := value; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-value]+1)
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

