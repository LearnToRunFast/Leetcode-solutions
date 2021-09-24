/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (39.00%)
 * Likes:    579
 * Dislikes: 0
 * Total Accepted:    89K
 * Total Submissions: 227.5K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 0
 * 0
 *
 *
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(k int, prices []int) int {
	if k < 1 {
		return 0
	}
	dp := make([][2]int, k+1)
	// 0 is buy 1 is sell
	for i := 0; i <= k; i++ {
		dp[i][0] = math.MinInt64
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j-1][1]-prices[i])
			dp[j][1] = max(dp[j][1], dp[j][0]+prices[i])
		}
	}
	return dp[k][1]

}

// @lc code=end

