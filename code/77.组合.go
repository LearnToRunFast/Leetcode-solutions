/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (76.53%)
 * Likes:    508
 * Dislikes: 0
 * Total Accepted:    137.2K
 * Total Submissions: 179.2K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */
// @lc code=start
func backtracking(path []int, n, curr, k int, res *[][]int) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	if n-curr+1 < k-len(path) {
		return
	}
	for i := curr; i <= n; i++ {
		path = append(path, i)
		backtracking(path, n, i+1, k, res)
		path = path[:len(path)-1]
	}
}
func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	curr := 1
	backtracking(path, n, curr, k, &res)
	return res
}

// @lc code=end
