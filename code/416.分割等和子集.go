/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode-cn.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (50.84%)
 * Likes:    1006
 * Dislikes: 0
 * Total Accepted:    184.2K
 * Total Submissions: 362.1K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,11,5]
 * 输出：true
 * 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,5]
 * 输出：false
 * 解释：数组不能分割成两个元素和相等的子集。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// define dp[i][j] to be true if 0 - i has j sum.
	// relation if j≥nums[i]: dp[i][j] = dp[i-1][j] if dont pick nums[i]
	// or dp[i-1][j-nums[i]] pick nums[i]
	// if j < nums[i]: dp[i][j] = dp[i-1][j], dont pick

	// observer that dp[i] can be constructed from dp[i-1]
	// so we can use one dimentional array to save space
	// if we loop through the array in increase order, we will update the dp[i-v] early
	// so that it can affect the result of dp[i]
	// so we need to decrease from target to v
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] = dp[i] || dp[i-v]
		}
	}
	return dp[target]
}

// @lc code=end

