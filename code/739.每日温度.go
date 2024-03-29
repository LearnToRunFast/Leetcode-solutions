/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (67.85%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    211.8K
 * Total Submissions: 311.6K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
 *
 * 示例 1:
 *
 *
 * 输入: temperatures = [73,74,75,71,69,72,76,73]
 * 输出: [1,1,4,2,1,1,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: temperatures = [30,40,50,60]
 * 输出: [1,1,1,0]
 *
 *
 * 示例 3:
 *
 *
 * 输入: temperatures = [30,60,90]
 * 输出: [1,1,0]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 30
 *
 *
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i, t := range temperatures {
		lastIdx := len(stack) - 1
		for len(stack) > 0 && t > temperatures[stack[lastIdx]] {
			j := stack[lastIdx]
			res[j] = i - j
			stack = stack[:lastIdx]
			lastIdx--
		}
		stack = append(stack, i)
	}
	return res
}

// @lc code=end

