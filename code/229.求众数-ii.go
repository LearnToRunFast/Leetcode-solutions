/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 *
 * https://leetcode-cn.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (45.39%)
 * Likes:    354
 * Dislikes: 0
 * Total Accepted:    28.6K
 * Total Submissions: 62.9K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
 *
 * 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[3,2,3]
 * 输出：[3]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：[1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：[1,1,1,3,3,2,2,2]
 * 输出：[1,2]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^9
 *
 *
 */

// @lc code=start
func majorityElement(nums []int) []int {
	ans := make([]int, 0, 2)
	n := len(nums)
	if nums == nil || n == 0 {
		return ans
	}
	cand1, cand2 := nums[0], nums[0]
	count1, count2 := 0, 0
	for _, v := range nums {
		switch v {
		case cand1:
			count1++
		case cand2:
			count2++
		default:
			if count1 == 0 {
				cand1 = v
				count1++
				continue
			}
			if count2 == 0 {
				cand2 = v
				count2++
				continue
			}
			count1--
			count2--
		}
	}
	count1 = 0
	count2 = 0
	for _, v := range nums {
		switch v {
		case cand1:
			count1++
		case cand2:
			count2++
		}
	}
	target := n / 3
	if count1 > target {
		ans = append(ans, cand1)
	}
	if count2 > target {
		ans = append(ans, cand2)
	}
	return ans

}

// @lc code=end

