/*
 * @lc app=leetcode.cn id=282 lang=golang
 *
 * [282] 给表达式添加运算符
 *
 * https://leetcode-cn.com/problems/expression-add-operators/description/
 *
 * algorithms
 * Hard (37.16%)
 * Likes:    209
 * Dislikes: 0
 * Total Accepted:    5.8K
 * Total Submissions: 15.3K
 * Testcase Example:  '"123"\n6'
 *
 * 给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = "123", target = 6
 * 输出: ["1+2+3", "1*2*3"]
 *
 *
 * 示例 2:
 *
 *
 * 输入: num = "232", target = 8
 * 输出: ["2*3+2", "2+3*2"]
 *
 * 示例 3:
 *
 *
 * 输入: num = "105", target = 5
 * 输出: ["1*0+5","10-5"]
 *
 * 示例 4:
 *
 *
 * 输入: num = "00", target = 0
 * 输出: ["0+0", "0-0", "0*0"]
 *
 *
 * 示例 5:
 *
 *
 * 输入: num = "3456237490", target = 9191
 * 输出: []
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * num 仅含数字
 *
 *
 */

// @lc code=start
func addOperators(num string, target int) []string {
	ans := []string{}

	var backtracking func(num string, prefix *[]string, sum, prevVal int)
	backtracking = func(num string, prefix *[]string, sum, prevVal int) {
		n := len(num)
		if n == 0 {
			if sum == target {
				ans = append(ans, strings.Join(*prefix, ""))
			}
			return
		}
		for i := 1; i <= n; i++ {
			if num[0] == '0' && i > 1 {
				break
			}
			curr, _ := strconv.Atoi(num[:i])
			if len(*prefix) == 0 {
				*prefix = append(*prefix, num[:i])
				backtracking(num[i:], prefix, curr, curr)
				*prefix = (*prefix)[:len(*prefix)-1]
				continue
			}
			*prefix = append(*prefix, "+")
			*prefix = append(*prefix, num[:i])
			backtracking(num[i:], prefix, sum+curr, curr)
			*prefix = (*prefix)[:len(*prefix)-2]
			*prefix = append(*prefix, "-")
			*prefix = append(*prefix, num[:i])
			backtracking(num[i:], prefix, sum-curr, -curr)
			*prefix = (*prefix)[:len(*prefix)-2]
			*prefix = append(*prefix, "*")
			*prefix = append(*prefix, num[:i])
			backtracking(num[i:], prefix, sum-prevVal+curr*prevVal, curr*prevVal)
			*prefix = (*prefix)[:len(*prefix)-2]

		}

	}
	backtracking(num, &[]string{}, 0, 0)
	return ans
}

// @lc code=end

