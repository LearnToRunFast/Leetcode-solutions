/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (55.07%)
 * Likes:    232
 * Dislikes: 0
 * Total Accepted:    24K
 * Total Submissions: 43.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的 非空 整数数组，每次操作将会使 n - 1 个元素增加 1。找出让数组所有元素相等的最小操作次数。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * [1,2,3]
 * 输出：
 * 3
 * 解释：
 * 只需要3次操作（注意每次操作会增加两个元素的值）：
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 *
 *
 */

// @lc code=start
func minMoves(nums []int) int {
	moves := 0
	min := 1<<31 - 1
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		moves += num - min
	}
	return moves
}

// @lc code=end

