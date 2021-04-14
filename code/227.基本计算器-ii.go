/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (43.25%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    66K
 * Total Submissions: 152.5K
 * Testcase Example:  '"3+2*2"'
 *
 * 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
 *
 * 整数除法仅保留整数部分。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3+2*2"
 * 输出：7
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = " 3/2 "
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = " 3+5 / 2 "
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
 * s 表示一个 有效表达式
 * 表达式中的所有整数都是非负整数，且在范围 [0, 2^31 - 1] 内
 * 题目数据保证答案是一个 32-bit 整数
 *
 *
 *
 *
 */

// @lc code=start
func cal(ans, curr, prev *int, op rune) {
	switch op {
	case '+':
		*ans += *prev
		*prev = *curr
	case '-':
		*ans += *prev
		*prev = -(*curr)
	case '*':
		*prev *= *curr
	case '/':
		*prev /= *curr
	}
}
func calculate(s string) int {
	ans := 0
	curr, prev := 0, 0
	op := '+'
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if c >= '0' && c <= '9' {
			curr = curr*10 + int(c-'0')
			continue
		}
		cal(&ans, &curr, &prev, op)
		op = c
		curr = 0

	}
	cal(&ans, &curr, &prev, op)
	ans += prev
	return ans
}

// @lc code=end

