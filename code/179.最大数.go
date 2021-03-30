/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (37.95%)
 * Likes:    485
 * Dislikes: 0
 * Total Accepted:    54.7K
 * Total Submissions: 143.6K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数 nums，重新排列它们每个数字的顺序（每个数字不可拆分）使之组成一个最大的整数。
 *
 * 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,2]
 * 输出："210"
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出："1"
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [10]
 * 输出："10"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

package main

import (
	"sort"
	"strconv"
	"strings"
)

// @lc code=start
// type ByOrder []int

// func (b ByOrder) Len() int { return len(b) }
// func (b ByOrder) Less(i, j int) bool {
// 	t1 := strconv.Itoa(b[i])
// 	t2 := strconv.Itoa(b[j])
// 	x := t1 + t2
// 	y := t2 + t1
// 	return x > y
// }
// func (b ByOrder) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// func largestNumber(nums []int) string {
// 	if nums == nil || len(nums) == 0 {
// 		return ""
// 	}
// 	sort.Sort(ByOrder(nums))
// 	if nums[0] == 0 {
// 		return "0"
// 	}
// 	var ans strings.Builder
// 	for _, v := range nums {
// 		ans.WriteString(strconv.Itoa(v))
// 	}
// 	return ans.String()
// }
type ByOrder []string

func (b ByOrder) Len() int { return len(b) }
func (b ByOrder) Less(i, j int) bool {
	x := b[i] + b[j]
	y := b[j] + b[i]
	return y < x
}
func (b ByOrder) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, v := range nums {
		strNums[i] = strconv.Itoa(v)
	}
	sort.Sort(ByOrder(strNums))
	if strNums[0] == "0" {
		return "0"
	}
	var ans strings.Builder
	for _, v := range strNums {
		ans.WriteString(v)
	}
	return ans.String()
}

// @lc code=end
