/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 *
 * https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (71.07%)
 * Likes:    322
 * Dislikes: 0
 * Total Accepted:    174.8K
 * Total Submissions: 246.2K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 解释：平方后，数组变为 [16,1,0,9,100]
 * 排序后，数组变为 [0,1,9,16,100]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 10^4
 * -10^4
 * nums 已按 非递减顺序 排序
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 请你设计时间复杂度为 O(n) 的算法解决本问题
 *
 *
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j, k := 0, n-1, n-1
	for i <= j {
		ls := nums[i] * nums[i]
		rs := nums[j] * nums[j]
		if ls > rs {
			res[k] = ls
			i++
		} else {
			res[k] = rs
			j--
		}
		k--
	}
	return res

}

// @lc code=end

