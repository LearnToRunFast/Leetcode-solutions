/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 *
 * https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (28.73%)
 * Likes:    211
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 69.7K
 * Testcase Example:  '1\n2'
 *
 * 给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
 *
 * 如果小数部分为循环小数，则将循环的部分括在括号内。
 *
 * 如果存在多个答案，只需返回 任意一个 。
 *
 * 对于所有给定的输入，保证 答案字符串的长度小于 10^4 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numerator = 1, denominator = 2
 * 输出："0.5"
 *
 *
 * 示例 2：
 *
 *
 * 输入：numerator = 2, denominator = 1
 * 输出："2"
 *
 *
 * 示例 3：
 *
 *
 * 输入：numerator = 2, denominator = 3
 * 输出："0.(6)"
 *
 *
 * 示例 4：
 *
 *
 * 输入：numerator = 4, denominator = 333
 * 输出："0.(012)"
 *
 *
 * 示例 5：
 *
 *
 * 输入：numerator = 1, denominator = 5
 * 输出："0.2"
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31
 * denominator != 0
 *
 *
 */

// @lc code=start
package main

import "strconv"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func dTob(a int) byte {
	return byte('0' + a)
}
func fractionToDecimal(numerator int, denominator int) string {
	ans := []byte{}
	if numerator < 0 && denominator > 0 || (numerator > 0 && denominator < 0) {
		ans = append(ans, '-')
	}

	a, d := abs(numerator), abs(denominator)
	q, r := a/d, a%d
	ans = append(ans, []byte(strconv.FormatInt(int64(q), 10))...)
	if r == 0 {
		return string(ans)
	}
	//ans = append(ans, dTob(q))
	ans = append(ans, '.')
	pos, repeated := len(ans), -1
	lookup := map[int]int{r: pos}

	for r != 0 {
		q = r * 10
		q, r = q/d, q%d
		ans = append(ans, dTob(q))
		pos++
		if v, ok := lookup[r]; ok {
			repeated = v
			break
		}
		lookup[r] = pos

	}

	if repeated != -1 {
		n := len(ans)
		// repeat part
		repeat := make([]byte, n-repeated)
		copy(repeat, ans[repeated:n])
		ans[repeated] = '('
		ans = ans[:repeated+1]
		ans = append(ans, repeat...)
		ans = append(ans, ')')
	}
	return string(ans)
}

// @lc code=end
