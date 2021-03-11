/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (52.88%)
 * Likes:    710
 * Dislikes: 0
 * Total Accepted:    101.6K
 * Total Submissions: 191.9K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 *
 *
 *
 * 进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^9
 *
 *
 */

package main

// @lc code=start

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	max := 0
	basket := map[int]int{}

	for _, v := range nums {
		_, ok := basket[v]
		if !ok {
			left := basket[v-1]
			right := basket[v+1]
			basket[v] = left + right + 1
			if left > 0 {
				basket[v-left] = basket[v]
			}
			if right > 0 {
				basket[v+right] = basket[v]
			}
			if max < basket[v] {
				max = basket[v]
			}

		}
	}
	return max
}

// @lc code=end
