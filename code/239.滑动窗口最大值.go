/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (49.68%)
 * Likes:    1240
 * Dislikes: 0
 * Total Accepted:    202.1K
 * Total Submissions: 406.7K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,-1], k = 1
 * 输出：[1,-1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [9,11], k = 2
 * 输出：[11]
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums = [4,-2], k = 2
 * 输出：[4]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 * 1
 *
 *
 */

// @lc code=start
type Queue struct {
	data []int
}

func (q *Queue) Push(v int) {
	for q.Len() > 0 && q.data[q.Len()-1] < v {
		q.data = q.data[:q.Len()-1]
	}
	q.data = append(q.data, v)
}
func (q *Queue) Pop(v int) {
	if q.Len() > 0 && q.Peek() == v {
		q.data = q.data[1:]
	}
}
func (q *Queue) Len() int {
	return len(q.data)
}
func (q *Queue) Peek() int {
	if q.Len() == 0 {
		return 0
	}
	return q.data[0]
}
func maxSlidingWindow(nums []int, k int) []int {
	ansLen := len(nums) - k + 1
	ans := make([]int, ansLen)
	if len(nums) == 0 {
		return ans
	}
	q := &Queue{}
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ans[0] = q.Peek()
	for i := 1; i < ansLen; i++ {
		q.Pop(nums[i-1])
		q.Push(nums[i+k-1])
		ans[i] = q.Peek()
	}
	return ans
}

// @lc code=end

