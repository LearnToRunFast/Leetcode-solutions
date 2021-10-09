/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (58.14%)
 * Likes:    2748
 * Dislikes: 0
 * Total Accepted:    330.4K
 * Total Submissions: 567.9K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func trap(height []int) int {
// 	n := len(height)
// 	if n < 3 {
// 		return 0
// 	}
// 	maxLeft := make([]int, n)
// 	maxRight := make([]int, n)

// 	maxLeft[0] = height[0]
// 	for i := 1; i < n; i++ {
// 		maxLeft[i] = max(maxLeft[i-1], height[i])
// 	}
// 	maxRight[n-1] = height[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		maxRight[i] = max(maxRight[i+1], height[i])
// 	}

// 	ans := 0
// 	for i, h := range height {
// 		rain := min(maxLeft[i], maxRight[i]) - h
// 		if rain > 0 {
// 			ans += rain
// 		}
// 	}
// 	return ans
// }
func trap(heights []int) int {
	stack := make([]int, 0)
	ans := 0
	for i, h := range heights {
		for len(stack) > 0 && h > heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			currH := min(heights[left], h) - heights[mid]
			currW := i - left - 1
			ans += currH * currW
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

