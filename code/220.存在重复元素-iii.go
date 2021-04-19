/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (26.54%)
 * Likes:    311
 * Dislikes: 0
 * Total Accepted:    31.9K
 * Total Submissions: 120.1K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j
 * 的差的绝对值也小于等于 ķ 。
 *
 * 如果存在则返回 true，不存在返回 false。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3, t = 0
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1, t = 2
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出: false
 *
 */
package main

// @lc code=start
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 0 {
		return false
	}
	bucket := map[int]int{}
	dis := t + 1
	for i, num := range nums {
		pos := num / dis
		if num < 0 {
			pos--
		}
		if _, ok := bucket[pos]; ok {
			return true
		}
		if v, ok := bucket[pos-1]; ok && abs(num-v) < dis {
			return true
		}
		if v, ok := bucket[pos+1]; ok && abs(num-v) < dis {
			return true
		}
		bucket[pos] = num
		if i >= k {
			delete(bucket, nums[i-k]/dis)
		}
	}
	return false

}

// @lc code=end
