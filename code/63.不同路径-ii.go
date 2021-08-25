/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (38.83%)
 * Likes:    614
 * Dislikes: 0
 * Total Accepted:    171.6K
 * Total Submissions: 440.5K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
 * 输出：2
 * 解释：
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 * 示例 2：
 *
 *
 * 输入：obstacleGrid = [[0,1],[0,0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == obstacleGrid.length
 * n == obstacleGrid[i].length
 * 1
 * obstacleGrid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
func getSteps(obstacleGrid, dp [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(obstacleGrid) || j >= len(obstacleGrid[0]) {
		return 0
	}
	return dp[i][j]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, n)
	dp[0] = make([]int, m)
	for i := range dp[0] {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			if obstacleGrid[i][j] == 0 {
				top := getSteps(obstacleGrid, dp, i-1, j)
				left := getSteps(obstacleGrid, dp, i, j-1)
				dp[i][j] = top + left
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[n-1][m-1]
}

// @lc code=end

