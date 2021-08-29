/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 *
 * https://leetcode-cn.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (45.90%)
 * Likes:    478
 * Dislikes: 0
 * Total Accepted:    63.9K
 * Total Submissions: 139.2K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
 *
 * 两个相邻元素间的距离为 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
 * 输出：[[0,0,0],[0,1,0],[0,0,0]]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
 * 输出：[[0,0,0],[0,1,0],[1,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1
 * 1
 * mat[i][j] is either 0 or 1.
 * mat 中至少有一个 0
 *
 *
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
				continue
			}
			res[i][j] = 1<<31 - 1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				res[i][j] = min(res[i][j], res[i-1][j]+1)
			}
			if j-1 >= 0 {
				res[i][j] = min(res[i][j], res[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j+1 < n {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res
}

// @lc code=end

