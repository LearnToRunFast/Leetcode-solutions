/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode-cn.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (56.06%)
 * Likes:    375
 * Dislikes: 0
 * Total Accepted:    69.1K
 * Total Submissions: 123.3K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * 输出:
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * 输出:
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 *
 * 进阶:
 *
 *
 * 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个常数空间的解决方案吗？
 *
 *
 */
package main

// @lc code=start
func markCol(matrix [][]int, j int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][j] = 0
	}
}
func markRow(matrix [][]int, j int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[j][i] = 0
	}
}
func mark(matrix [][]int, i, j int) {
	matrix[i][0] = 0
	matrix[0][j] = 0
}
func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	setFirstCol, setFirstRow := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			setFirstCol = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			setFirstRow = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mark(matrix, i, j)
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			markRow(matrix, i)
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			markCol(matrix, j)
		}
	}
	if setFirstCol {
		markCol(matrix, 0)
	}
	if setFirstRow {
		markRow(matrix, 0)
	}
}

// @lc code=end
