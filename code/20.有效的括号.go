/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.50%)
 * Likes:    2740
 * Dislikes: 0
 * Total Accepted:    819.3K
 * Total Submissions: 1.8M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "([)]"
 * 输出：false
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "{[]}"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	lookup := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 || stack[len(stack)-1] != lookup[c] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// @lc code=end

