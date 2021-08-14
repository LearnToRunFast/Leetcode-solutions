/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (65.76%)
 * Likes:    1082
 * Dislikes: 0
 * Total Accepted:    290.6K
 * Total Submissions: 441.4K
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 3, n = 7
 * 输出：28
 *
 * 示例 2：
 *
 *
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 *
 *
 * 示例 3：
 *
 *
 * 输入：m = 7, n = 3
 * 输出：28
 *
 *
 * 示例 4：
 *
 *
 * 输入：m = 3, n = 3
 * 输出：6
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 题目数据保证答案小于等于 2 * 10^9
 *
 *
 */

// @lc code=start
func getPaths(i, j, m, n int, dp [][]int) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	return dp[i][j]
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		if i != 0 {
			dp[i] = make([]int, n)
			for j := 0; j < n; j++ {
				up := getPaths(i-1, j, m, n, dp)
				left := getPaths(i, j-1, m, n, dp)
				dp[i][j] = up + left
			}
		}
	}
	return dp[m-1][n-1]

}

// @lc code=end

