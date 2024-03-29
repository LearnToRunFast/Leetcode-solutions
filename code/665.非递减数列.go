/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 *
 * https://leetcode-cn.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Medium (27.09%)
 * Likes:    633
 * Dislikes: 0
 * Total Accepted:    77K
 * Total Submissions: 284K
 * Testcase Example:  '[4,2,3]'
 *
 * 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
 *
 * 我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 ，总满足 nums[i] 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [4,2,3]
 * 输出: true
 * 解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [4,2,1]
 * 输出: false
 * 解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * - 10 ^ 5
 *
 *
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x > y {
			count++
			if count > 1 {
				return false
			}
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}

// @lc code=end

