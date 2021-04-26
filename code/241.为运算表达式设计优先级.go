/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 *
 * https://leetcode-cn.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (73.39%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    23.4K
 * Total Submissions: 31.8K
 * Testcase Example:  '"2-1-1"'
 *
 * 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及
 * * 。
 *
 * 示例 1:
 *
 * 输入: "2-1-1"
 * 输出: [0, 2]
 * 解释:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * 示例 2:
 *
 * 输入: "2*3-4*5"
 * 输出: [-34, -14, -10, -10, 10]
 * 解释:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 */

// @lc code=start

func diffWaysToCompute(expression string) []int {
	num, err := strconv.Atoi(expression)
	if err == nil {
		return []int{num}
	}
	ans := []int{}
	for i, c := range expression {
		switch c {
		case '+', '-', '*':
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, num1 := range left {
				for _, num2 := range right {
					res := 0
					switch c {
					case '+':
						res = num1 + num2
					case '-':
						res = num1 - num2
					case '*':
						res = num1 * num2
					}
					ans = append(ans, res)
				}
			}
		}
	}
	return ans
}

// @lc code=end

