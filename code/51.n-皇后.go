/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (73.50%)
 * Likes:    756
 * Dislikes: 0
 * Total Accepted:    104.1K
 * Total Submissions: 141.3K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 *
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[["Q"]]
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

func generateBoard(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		qPosition := queens[i]
		row[qPosition] = 'Q'
		board[i] = string(row)
	}
	return board
}
func backtrack(queens []int, ans *[][]string, n, startRow, colSlots, backSlashSlots, forwardSlashSlots int) {
	if startRow == n {
		*ans = append(*ans, generateBoard(queens, n))
		return
	}
	mask := (1 << uint(n)) - 1
	slots := colSlots | backSlashSlots | forwardSlashSlots // 1 means occupied
	availableSlots := mask & (^slots)                      // 1 means available
	for availableSlots != 0 {
		position := availableSlots & (-availableSlots)         // get the rightmost 1
		availableSlots = availableSlots & (availableSlots - 1) // remove the rightmost 1
		countPosition := bits.OnesCount(uint(position - 1))
		queens[startRow] = countPosition
		newColSlots := colSlots | position
		newBackSlots := (backSlashSlots | position) << 1
		newForwardSlots := (forwardSlashSlots | position) >> 1
		backtrack(queens, ans, n, startRow+1, newColSlots, newBackSlots, newForwardSlots)
		queens[startRow] = 0
	}

}
func solveNQueens(n int) [][]string {
	ans := [][]string{}
	// one dimension array for position of queen in each row
	queens := make([]int, n)

	startRow, colSlots, backSlashSlots, forwardSlashSlots := 0, 0, 0, 0
	backtrack(queens, &ans, n, startRow, colSlots, backSlashSlots, forwardSlashSlots)
	return ans
}

// @lc code=end

