/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (60.58%)
 * Likes:    902
 * Dislikes: 0
 * Total Accepted:    116.1K
 * Total Submissions: 191.2K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
 *
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 *
 *
 * 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 *
 *
 * 示例:
 *
 * 输入: [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 *
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// 0 hold
	// 1 sold in freeze (sold at that day)
	// 2 sold no in freeze
	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		// yesterday alr hold or sold no in freeze
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}

	return max(dp[n-1][1], dp[n-1][2])
}

// @lc code=end

