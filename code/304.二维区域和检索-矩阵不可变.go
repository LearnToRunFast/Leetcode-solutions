/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-2d-immutable/description/
 *
 * algorithms
 * Medium (46.94%)
 * Likes:    140
 * Dislikes: 0
 * Total Accepted:    15.2K
 * Total Submissions: 32.2K
 * Testcase Example:  '["NumMatrix","sumRegion","sumRegion","sumRegion"]\n' +
  '[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]'
 *
 * 给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
 *
 *
 * 上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
 *
 * 示例:
 *
 * 给定 matrix = [
 * ⁠ [3, 0, 1, 4, 2],
 * ⁠ [5, 6, 3, 2, 1],
 * ⁠ [1, 2, 0, 1, 5],
 * ⁠ [4, 1, 0, 1, 7],
 * ⁠ [1, 0, 3, 0, 5]
 * ]
 *
 * sumRegion(2, 1, 4, 3) -> 8
 * sumRegion(1, 1, 2, 2) -> 11
 * sumRegion(1, 2, 2, 4) -> 12
 *
 *
 * 说明:
 *
 *
 * 你可以假设矩阵不可变。
 * 会多次调用 sumRegion 方法。
 * 你可以假设 row1 ≤ row2 且 col1 ≤ col2。
 *
 *
*/

// @lc code=start
type NumMatrix struct {
	sum *[][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m)
	prevRow := make([]int, n+1)
	for i := range matrix {
		sum[i] = make([]int, n)
		for j, v := range matrix[i] {
			prev := 0
			if j > 0 {
				prev = sum[i][j-1]
			}
			sum[i][j] = v + prev + prevRow[j+1] - prevRow[j]
		}
		// update preRow
		for k := 1; k <= n; k++ {
			prevRow[k] = sum[i][k-1]
		}
	}
	return NumMatrix{&sum}
}
func getValue(sum *[][]int, i, j, m, n int) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	return (*sum)[i][j]
}
func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	matrix := *this.sum
	m, n := len(matrix), len(matrix[0])
	sum := matrix[r2][c2]
	prevTop := getValue(this.sum, r1-1, c2, m, n)
	prevLeft := getValue(this.sum, r2, c1-1, m, n)
	prevTopLeft := getValue(this.sum, r1-1, c1-1, m, n)
	ans := sum - prevTop - prevLeft + prevTopLeft
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

