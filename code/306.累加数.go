/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 *
 * https://leetcode-cn.com/problems/additive-number/description/
 *
 * algorithms
 * Medium (33.34%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    13.3K
 * Total Submissions: 40K
 * Testcase Example:  '"112358"'
 *
 * 累加数是一个字符串，组成它的数字可以形成累加序列。
 *
 * 一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
 *
 * 给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。
 *
 * 说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
 *
 * 示例 1:
 *
 * 输入: "112358"
 * 输出: true
 * 解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
 *
 *
 * 示例 2:
 *
 * 输入: "199100199"
 * 输出: true
 * 解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
 *
 * 进阶:
 * 你如何处理一个溢出的过大的整数输入?
 *
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	n := len(num)
	if n < 3 {
		return false
	}
	var backtracking func(sum, prev int64, currIdx, count int) bool
	backtracking = func(sum, prev int64, currIdx, count int) bool {
		if currIdx == n {
			return count >= 3
		}
		var val int64
		for i := currIdx; i < n; i++ {
			if i > currIdx && num[currIdx] == '0' {
				break
			}
			val = int64(num[i]-'0') + val*10

			if count >= 2 {
				if val < sum {
					continue
				} else if val > sum {
					break
				}
			}
			isValid := backtracking(prev+val, val, i+1, count+1)
			if isValid {
				return true
			}
		}
		return false
	}
	return backtracking(0, 0, 0, 0)
}

// @lc code=end

