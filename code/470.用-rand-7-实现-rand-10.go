/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 *
 * https://leetcode-cn.com/problems/implement-rand10-using-rand7/description/
 *
 * algorithms
 * Medium (55.30%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    30.1K
 * Total Submissions: 54.5K
 * Testcase Example:  '1'
 *
 * 已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
 *
 * 不要使用系统的 Math.random() 方法。
 *
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: 1
 * 输出: [7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: 2
 * 输出: [8,4]
 *
 *
 * 示例 3:
 *
 *
 * 输入: 3
 * 输出: [8,1,10]
 *
 *
 *
 *
 * 提示:
 *
 *
 * rand7 已定义。
 * 传入参数: n 表示 rand10 的调用次数。
 *
 *
 *
 *
 * 进阶:
 *
 *
 * rand7()调用次数的 期望值 是多少 ?
 * 你能否尽量少调用 rand7() ?
 *
 *
 */

// @lc code=start
func rand10() int {
	var x int
	var y int
	for {
		x = rand7()
		y = rand7()
		val := y + (x-1)*7
		if val <= 40 {
			return val%10 + 1
		}
		x = val - 40
		y = rand7()
		val = y + (x-1)*7
		if val <= 60 {
			return val%10 + 1
		}
		x = val - 60
		y = rand7()
		val = y + (x-1)*7
		if val <= 20 {
			return val%10 + 1
		}
	}
}

// @lc code=end

