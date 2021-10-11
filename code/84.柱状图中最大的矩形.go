/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (43.34%)
 * Likes:    1581
 * Dislikes: 0
 * Total Accepted:    185.3K
 * Total Submissions: 427.4K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入： heights = [2,4]
 * 输出： 4
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
func largestRectangleArea(heights []int) int {
	heights = append([]int{-1}, heights...) // head guard
	heights = append(heights, -1)           // tail guard
	stack := make([]int, 0)
	ans := 0
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[mid]
			width := i - stack[len(stack)-1] - 1
			ans = max(ans, height*width)
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

