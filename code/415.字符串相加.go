/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (52.90%)
 * Likes:    383
 * Dislikes: 0
 * Total Accepted:    122.2K
 * Total Submissions: 230.9K
 * Testcase Example:  '"11"\n"123"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 *
 *
 * 提示：
 *
 *
 * num1 和num2 的长度都小于 5100
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
 *
 *
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	carry := byte(0)
	max := n
	if max < m {
		max = m
	}
	ans := make([]byte, max+1)
	j := max
	for i := 0; i < max; i++ {
		a, b := byte(48), byte(48)

		if i < n {
			a = num1[n-1-i]
		}
		if i < m {
			b = num2[m-1-i]
		}
		a = a + b - '0' + carry
		if a > '9' {
			carry = 1
			a = a - 10
		} else {
			carry = 0
		}
		ans[j] = a
		j--
	}
	if carry > 0 {
		ans[0] = '1'
	} else {
		ans = ans[1:]
	}
	return string(ans)
}

// @lc code=end

