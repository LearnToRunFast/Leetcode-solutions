/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (52.46%)
 * Likes:    334
 * Dislikes: 0
 * Total Accepted:    82.9K
 * Total Submissions: 158.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：24
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [-1,-2,-3]
 * 输出：-6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3
 * -1000
 *
 *
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int) int {
	max1, max2, max3 := -1<<63, -1<<63, -1<<63
	min1, min2 := 1<<63-1, 1<<63-1
	for _, n := range nums {
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}
	}
	// if all positive, max1 * max2 * max3
	// if all negative, max1 * max2 * max3
	// if mix, either max1 * max2 * max3 or min1 * min2 * max1
	return max(max1*max2*max3, min1*min2*max1)
}

// @lc code=end

