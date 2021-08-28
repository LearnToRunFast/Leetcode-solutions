/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 *
 * https://leetcode-cn.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (45.44%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    41K
 * Total Submissions: 89K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 *
 *
 *
 * 示例:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * 输出:  [1,2,4,7,5,3,6,8,9]
 *
 * 解释:
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 给定矩阵中的元素总数不会超过 100000 。
 *
 *
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])
	res := make([]int, 0, m*n)
	isUp := true
	i, j := 0, 0
	for i < m && j < n {
		res = append(res, mat[i][j])
		if isUp {
			if i != 0 && j != n-1 {
				i--
				j++
			} else {
				isUp = false
				if j == n-1 {
					i++
				} else {
					j++
				}
			}
			continue
		}
		if i != m-1 && j != 0 {
			i++
			j--
		} else {
			isUp = true
			if i == m-1 {
				j++
			} else {
				i++
			}
		}
	}

	return res
}

// @lc code=end

