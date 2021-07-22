/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132 模式
 *
 * https://leetcode-cn.com/problems/132-pattern/description/
 *
 * algorithms
 * Medium (36.04%)
 * Likes:    529
 * Dislikes: 0
 * Total Accepted:    52.4K
 * Total Submissions: 145.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k]
 * 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
 *
 * 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：false
 * 解释：序列中不存在 132 模式的子序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,1,4,2]
 * 输出：true
 * 解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [-1,3,2,0]
 * 输出：true
 * 解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1
 * -10^9
 *
 *
 */

// @lc code=start
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	n := len(nums)
	stack := make([]int, 0)
	k := -1 << 31
	for i := n - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			if k < stack[len(stack)-1] {
				k = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false
}

// @lc code=end

