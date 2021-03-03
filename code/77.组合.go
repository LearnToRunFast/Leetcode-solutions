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
package main

// @lc code=start

func combine(n int, k int) [][]int {
	if k <= 0 || k > n {
		return [][]int{}
	}
	perms := make([][]int, 0, n)
	perm := make([]int, 0, k)
	var backtracking func(start int)

	backtracking = func(start int) {
		if len(perm) == k {
			newPerm := append([]int{}, perm...)
			perms = append(perms, newPerm)
			return
		}
		end := n - (k - len(perm)) + 1
		for i := start; i <= end; i++ {
			perm = append(perm, i)
			backtracking(i + 1)
			perm = perm[:len(perm)-1]
		}
	}
	backtracking(1)
	return perms
}

// @lc code=end
