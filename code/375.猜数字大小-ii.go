/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 *
 * https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/description/
 *
 * algorithms
 * Medium (45.09%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    12.6K
 * Total Submissions: 27.9K
 * Testcase Example:  '10'
 *
 * 我们正在玩一个猜数游戏，游戏规则如下：
 *
 * 我从 1 到 n 之间选择一个数字，你来猜我选了哪个数字。
 *
 * 每次你猜错了，我都会告诉你，我选的数字比你的大了或者小了。
 *
 * 然而，当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。直到你猜到我选的数字，你才算赢得了这个游戏。
 *
 * 示例:
 *
 * n = 10, 我选择了8.
 *
 * 第一轮: 你猜我选择的数字是5，我会告诉你，我的数字更大一些，然后你需要支付5块。
 * 第二轮: 你猜是7，我告诉你，我的数字更大一些，你支付7块。
 * 第三轮: 你猜是9，我告诉你，我的数字更小一些，你支付9块。
 *
 * 游戏结束。8 就是我选的数字。
 *
 * 你最终要支付 5 + 7 + 9 = 21 块钱。
 *
 *
 * 给定 n ≥ 1，计算你至少需要拥有多少现金才能确保你能赢得这个游戏。
 *
 */

// @lc code=start
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func getMoneyAmount(n int) int {

	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 2; j <= n; j++ {
		for i := j - 1; i >= 1; i-- {
			if i == j {
				continue
			}
			val := 1<<31 - 1
			for k := i + 1; k <= j-1; k++ {
				val = min(val, k+max(dp[i][k-1], dp[k+1][j]))
			}
			dp[i][j] = min(val, min(i+dp[i+1][j], j+dp[i][j-1]))
		}
	}
	return dp[1][n]

}

// @lc code=end

