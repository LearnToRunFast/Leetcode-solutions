/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 *
 * https://leetcode-cn.com/problems/elimination-game/description/
 *
 * algorithms
 * Medium (46.25%)
 * Likes:    115
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 13.2K
 * Testcase Example:  '9'
 *
 * 给定一个从1 到 n 排序的整数列表。
 * 首先，从左到右，从第一个数字开始，每隔一个数字进行删除，直到列表的末尾。
 * 第二步，在剩下的数字中，从右到左，从倒数第一个数字开始，每隔一个数字进行删除，直到列表开头。
 * 我们不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
 * 返回长度为 n 的列表中，最后剩下的数字。
 *
 * 示例：
 *
 *
 * 输入:
 * n = 9,
 * 1 2 3 4 5 6 7 8 9
 * 2 4 6 8
 * 2 6
 * 6
 *
 * 输出:
 * 6
 *
 */

// @lc code=start
func lastRemaining(n int) int {
	left := 1
	step := 1
	toRight := true
	// 偶数数组：
	// 左->右  1 2 3 4 5 6 7 8
	// 		  2   4   6   8
	// 右->左  1 2 3 4 5 6 7 8
	// 		1   3   5   7
	// 奇数数组：
	// 左->右  1 2 3 4 5 6 7
	// 		  2   4   6
	// 右->左  1 2 3 4 5 6 7
	// 		  2   4   6
	for n > 1 {
		// update left
		if toRight || n%2 == 1 {
			left += step
		}
		n /= 2
		step *= 2
		toRight = !toRight
	}
	return left
}

// @lc code=end

