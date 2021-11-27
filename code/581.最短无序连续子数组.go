/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 *
 * https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Medium (40.53%)
 * Likes:    736
 * Dislikes: 0
 * Total Accepted:    103.6K
 * Total Submissions: 255.7K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
 *
 * 请你找出符合题意的 最短 子数组，并输出它的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,6,4,8,10,9,15]
 * 输出：5
 * 解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^5
 *
 *
 *
 *
 * 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
 *
 *
 *
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		if maxn > num {
			right = i
		} else {
			maxn = num
		}
		j := n - i - 1
		if minn < nums[j] {
			left = j
		} else {
			minn = nums[j]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

// @lc code=end

