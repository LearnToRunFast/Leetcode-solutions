/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 *
 * https://leetcode-cn.com/problems/degree-of-an-array/description/
 *
 * algorithms
 * Easy (60.46%)
 * Likes:    382
 * Dislikes: 0
 * Total Accepted:    68.6K
 * Total Submissions: 113.6K
 * Testcase Example:  '[1,2,2,3,1]'
 *
 * 给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。
 *
 * 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[1, 2, 2, 3, 1]
 * 输出：2
 * 解释：
 * 输入数组的度是2，因为元素1和2的出现频数最大，均为2.
 * 连续子数组里面拥有相同度的有如下所示:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * 最短连续子数组[2, 2]的长度为2，所以返回2.
 *
 *
 * 示例 2：
 *
 *
 * 输入：[1,2,2,3,1,4,2]
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums.length 在1到 50,000 区间范围内。
 * nums[i] 是一个在 0 到 49,999 范围内的整数。
 *
 *
 */

// @lc code=start
type Entry struct {
	count int
	left  int
	right int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findShortestSubArray(nums []int) int {
	entryMap := make(map[int]*Entry)
	for i, v := range nums {
		if _, ok := entryMap[v]; !ok {
			entryMap[v] = &Entry{
				count: 1,
				left:  i,
				right: i,
			}
		} else {
			entryMap[v].count++
			entryMap[v].right = i
		}
	}
	ans, maxCount := 0, 0
	for _, v := range entryMap {
		if v.count > maxCount {
			maxCount = v.count
			ans = v.right - v.left + 1
		} else if v.count == maxCount {
			ans = min(ans, v.right-v.left+1)
		}
	}
	return ans
}

// @lc code=end

