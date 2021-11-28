/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.62%)
 * Likes:    1193
 * Dislikes: 0
 * Total Accepted:    163.1K
 * Total Submissions: 365.8K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	preCount := map[int]int{0: 1}
	pre := 0
	count := 0
	for _, num := range nums {
		pre += num
		if preCount[pre-k] > 0 {
			count += preCount[pre-k]
		}
		preCount[pre]++
	}
	return count
}

// @lc code=end

