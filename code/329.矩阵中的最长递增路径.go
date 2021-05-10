/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 *
 * https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/description/
 *
 * algorithms
 * Hard (47.28%)
 * Likes:    463
 * Dislikes: 0
 * Total Accepted:    45.4K
 * Total Submissions: 96K
 * Testcase Example:  '[[9,9,4],[6,6,8],[2,1,1]]'
 *
 * 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
 *
 * 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
 * 输出：4
 * 解释：最长递增路径为 [1, 2, 6, 9]。
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
 * 输出：4
 * 解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [[1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * 0
 *
 *
 */

// @lc code=start
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return 0
	}
	// dir matrix
	dirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}

	// memo init
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	// dfs function
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}

		val := 1
		for _, dir := range dirs {
			x := i + dir[0]
			y := j + dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[i][j] < matrix[x][y] {
				val = max(val, dfs(x, y)+1)
			}
		}
		memo[i][j] = val
		return val
	}

	// main loop
	ans := 0
	for i, rows := range matrix {
		for j := range rows {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans

}

// @lc code=end

