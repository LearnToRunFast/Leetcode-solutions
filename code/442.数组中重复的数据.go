/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 *
 * https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/description/
 *
 * algorithms
 * Medium (69.24%)
 * Likes:    395
 * Dislikes: 0
 * Total Accepted:    38.1K
 * Total Submissions: 55K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 *
 * 找到所有出现两次的元素。
 *
 * 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 *
 * 示例：
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [2,3]
 *
 *
 */

// @lc code=start
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func findDuplicates(nums []int) []int {
	var res []int
	for i := range nums {
		num := abs(nums[i])
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] *= -1
		}
	}
	return res
}

// @lc code=end

