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
func drawBoard(queens *[]int, n int) []string {
	board := make([]string, 0)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[(*queens)[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// // tld is top left to bottom right diagonal
// // trd is top right to bottom left diagonal
func backtrack(queens *[]int, ans *[][]string, n, currRow, col, tld, trd int) {
	if currRow == n {
		*ans = append(*ans, drawBoard(queens, n))
		return
	}
	mask := (1 << n) - 1             // all 1
	takenSlots := ^(col | tld | trd) // invert all 1 to 0
	slots := mask & takenSlots       // find available slots
	for slots != 0 {
		pos := slots & (-slots)                            // get lowest bit of 1
		slots = slots & (slots - 1)                        // set lowest bit of 1 to 0
		(*queens)[currRow] = bits.OnesCount(uint(pos - 1)) // get column number by counting the number of 1
		backtrack(queens, ans, n, currRow+1, col|pos, (tld|pos)<<1, (trd|pos)>>1)
	}
}

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	queens := make([]int, n)
	backtrack(&queens, &ans, n, 0, 0, 0, 0)
	return ans
}

// @lc code=end

