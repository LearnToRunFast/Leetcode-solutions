/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.16%)
 * Likes:    236
 * Dislikes: 0
 * Total Accepted:    57K
 * Total Submissions: 69.4K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
 *
 *
 *
 *
 */

// @lc code=start
func backtrack(ans *int, row, n, col, d1, d2 int) {
	if row == n {
		(*ans)++
		return
	}

	mask := (1 << n) - 1
	slots := ^(col | d1 | d2)
	slots = mask & slots

	for slots != 0 {
		// -x = ~x + 1 => x & (-x) =>
		pos := slots & (-slots)     // get lowest bit-1
		slots = slots & (slots - 1) // set lowest bit-1 to 0
		backtrack(ans, row+1, n, col|pos, (d1|pos)<<1, (d2|pos)>>1)

	}
}
func totalNQueens(n int) int {
	ans := 0
	backtrack(&ans, 0, n, 0, 0, 0)
	return ans

}

// @lc code=end

