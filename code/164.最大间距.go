/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 *
 * https://leetcode-cn.com/problems/maximum-gap/description/
 *
 * algorithms
 * Hard (61.18%)
 * Likes:    354
 * Dislikes: 0
 * Total Accepted:    49.6K
 * Total Submissions: 81.1K
 * Testcase Example:  '[3,6,9,1]'
 *
 * 给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
 *
 * 如果数组元素个数小于 2，则返回 0。
 *
 * 示例 1:
 *
 * 输入: [3,6,9,1]
 * 输出: 3
 * 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
 *
 * 示例 2:
 *
 * 输入: [10]
 * 输出: 0
 * 解释: 数组元素个数小于 2，因此返回 0。
 *
 * 说明:
 *
 *
 * 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
 * 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
 *
 *
 */

// @lc code=start
package main

func radixSort(nums []int) []int {
	const BACKET_SIZE = 10
	n := len(nums)
	newArr := make([]int, n, n)
	currVal := 1
	var getMax func(nums []int) int
	getMax = func(nums []int) int {
		max := 0
		for _, v := range nums {
			if max < v {
				max = v
			}
		}
		return max
	}
	maxValue := getMax(nums)

	for ; maxValue >= currVal; currVal *= 10 {
		backet := make([]int, BACKET_SIZE, BACKET_SIZE)
		for _, v := range nums {
			pos := (v / currVal) % BACKET_SIZE
			backet[pos]++
		}
		for i := 1; i < BACKET_SIZE; i++ {
			backet[i] += backet[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			v := nums[i]
			pos := (v / currVal) % BACKET_SIZE
			backet[pos]--
			newArr[backet[pos]] = v
		}
		copy(nums, newArr)
	}
	return nums
}
func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	nums = radixSort(nums)

	maxGap := 0
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if maxGap < diff {
			maxGap = diff
		}
	}
	return maxGap
}

// @lc code=end
