/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (36.12%)
 * Likes:    227
 * Dislikes: 0
 * Total Accepted:    52.8K
 * Total Submissions: 146.2K
 * Testcase Example:  '[3,2,1]'
 *
 * 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[3, 2, 1]
 * 输出：1
 * 解释：第三大的数是 1 。
 *
 * 示例 2：
 *
 *
 * 输入：[1, 2]
 * 输出：2
 * 解释：第三大的数不存在, 所以返回最大的数 2 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：[2, 2, 3, 1]
 * 输出：1
 * 解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
 * 此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度 O(n) 的解决方案吗？
 *
 */

// @lc code=start
func thirdMax(nums []int) int {
	min := -1 << 61
	a1, a2, a3 := min, min, min

	for _, num := range nums {
		if num == a1 || num == a2 || num == a3 {
			continue
		}
		if num > a1 {
			a1, a2, a3 = num, a1, a2
		} else if num > a2 {
			a2, a3 = num, a2
		} else if num > a3 {
			a3 = num
		}
	}
	if a3 == min {
		a3 = a1
	}
	return a3
}

// @lc code=end

