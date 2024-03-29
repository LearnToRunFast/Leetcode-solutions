/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (60.43%)
 * Likes:    254
 * Dislikes: 0
 * Total Accepted:    114.4K
 * Total Submissions: 189.3K
 * Testcase Example:  '[1,1,0,1,1,1]'
 *
 * 给定一个二进制数组， 计算其中最大连续 1 的个数。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：[1,1,0,1,1,1]
 * 输出：3
 * 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
 *
 *
 *
 *
 * 提示：
 *
 *
 * 输入的数组只包含 0 和 1 。
 * 输入数组的长度是正整数，且不超过 10,000。
 *
 *
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}

// @lc code=end

