/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (61.77%)
 * Likes:    404
 * Dislikes: 0
 * Total Accepted:    69.1K
 * Total Submissions: 111.7K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 */
package main

import "sort"

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	perms := [][]int{[]int{}}
	perm := []int{}
	var getPerm func(nums []int)
	getPerm = func(nums []int) {
		n := len(nums)
		for i := 0; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			perm = append(perm, nums[i])
			perms = append(perms, append([]int{}, perm...))
			getPerm(nums[i+1:])
			perm = perm[:len(perm)-1]
		}
	}
	getPerm(nums)
	return perms
}

// @lc code=end
