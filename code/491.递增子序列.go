/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (54.92%)
 * Likes:    313
 * Dislikes: 0
 * Total Accepted:    42.3K
 * Total Submissions: 77.3K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
var ans [][]int
var temp []int

func dfs(curr, prev int, nums []int) {
	if curr == len(nums) {
		if len(temp) >= 2 {
			copy := append([]int{}, temp...)
			ans = append(ans, copy)
		}
		return
	}
	if nums[curr] >= prev {
		temp = append(temp, nums[curr])
		dfs(curr+1, nums[curr], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[curr] != prev {
		dfs(curr+1, prev, nums)
	}
}
func findSubsequences(nums []int) [][]int {
	ans = make([][]int, 0)
	temp = make([]int, 0)
	dfs(0, -1<<31, nums)
	return ans
}

// @lc code=end

