/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 *
 * https://leetcode-cn.com/problems/count-numbers-with-unique-digits/description/
 *
 * algorithms
 * Medium (51.45%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    21.8K
 * Total Submissions: 42.2K
 * Testcase Example:  '2'
 *
 * 给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10^n 。
 *
 * 示例:
 *
 * 输入: 2
 * 输出: 91
 * 解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
 *
 *
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n+1)
	sum := 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]*10 + (9*int(math.Pow10(i-2))-dp[i-1])*(i-1)
		sum += dp[i]
	}

	return int(math.Pow10(n)) - sum
}

// @lc code=end

