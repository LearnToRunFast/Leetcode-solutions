/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (73.68%)
 * Likes:    384
 * Dislikes: 0
 * Total Accepted:    106.2K
 * Total Submissions: 144.2K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 *
 * 说明：
 *
 *
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 *
 *
 * 示例 2:
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */

// @lc code=start
func dfs(k, n, start int, path []int, res *[][]int) {
	if n < 0 || k < 0 {
		return
	}
	if k == 0 && n == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i <= 9; i++ {
		dfs(k-1, n-i, i+1, append(path, i), res)
	}
}
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	dfs(k, n, 1, []int{}, &res)
	return res
}

// @lc code=end

