/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode-cn.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (68.66%)
 * Likes:    522
 * Dislikes: 0
 * Total Accepted:    52.6K
 * Total Submissions: 76.8K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,3,2]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [0,1,0,1,0,1,99]
 * 输出: 99
 *
 */

// @lc code=start
func singleNumber(nums []int) int {
	once, twice := 0, 0
	for _, v := range nums {
		once = ^twice & (once ^ v)
		twice = ^once & (twice ^ v)
	}
	return once
}

// @lc code=end

