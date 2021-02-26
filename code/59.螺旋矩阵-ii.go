/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (78.53%)
 * Likes:    305
 * Dislikes: 0
 * Total Accepted:    61.7K
 * Total Submissions: 78.5K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	left, top, right, bottom := 0, 0, n-1, n-1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	num := 1
	for {
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}
		bottom--
		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			ans[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

// @lc code=end

