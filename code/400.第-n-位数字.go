/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第 N 位数字
 *
 * https://leetcode-cn.com/problems/nth-digit/description/
 *
 * algorithms
 * Medium (40.58%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    15.9K
 * Total Submissions: 39.1K
 * Testcase Example:  '3'
 *
 * 在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 位数字。
 *
 *
 *
 * 注意：n 是正数且在 32 位整数范围内（n < 2^31）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：3
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：11
 * 输出：0
 * 解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。
 *
 *
 */

// @lc code=start
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	base := 10
	count := 2
	next := 9 * base * count
	n -= 9
	for n > next {
		n -= next
		base *= 10
		count++
		next = 9 * base * count
	}
	q := n / count
	r := n % count
	num := q + base
	if r == 0 {
		num--
		r = count
	}
	numStr := strconv.Itoa(num)
	ans, _ := strconv.Atoi(string(numStr[r-1]))
	return ans
}

// @lc code=end

