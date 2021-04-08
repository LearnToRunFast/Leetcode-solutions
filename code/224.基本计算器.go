/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 *
 * https://leetcode-cn.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (41.63%)
 * Likes:    539
 * Dislikes: 0
 * Total Accepted:    56.1K
 * Total Submissions: 134.8K
 * Testcase Example:  '"1 + 1"'
 *
 * 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "1 + 1"
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = " 2-1 + 2 "
 * 输出：3
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(1+(4+5+2)-3)+(6+8)"
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
 * s 表示一个有效的表达式
 *
 *
 */

// @lc code=start
func calculate(s string) int {
	ops := []int{1}
	sign := 1
	ans := 0
	n := len(s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = ops[len(ops)-1]
		case '-':
			sign = -ops[len(ops)-1]
		case '(':
			ops = append(ops, sign)
		case ')':
			ops = ops[:len(ops)-1]
		default:
			num := 0
			for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			i--
			ans += sign * num
		}
	}
	return ans
}

// @lc code=end

