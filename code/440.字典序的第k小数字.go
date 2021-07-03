/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 *
 * https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (37.51%)
 * Likes:    220
 * Dislikes: 0
 * Total Accepted:    13.7K
 * Total Submissions: 36.6K
 * Testcase Example:  '13\n2'
 *
 * 给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
 *
 * 注意：1 ≤ k ≤ n ≤ 10^9。
 *
 * 示例 :
 *
 *
 * 输入:
 * n: 13   k: 2
 *
 * 输出:
 * 10
 *
 * 解释:
 * 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
 *
 *
 */

// @lc code=start
func getCount(prefix, n int) int {
	count := 0
	next := prefix + 1
	for i := prefix; i <= n; i *= 10 {
		min := next
		if min > n+1 {
			min = n + 1
		}
		count += min - i
		next *= 10
	}
	return count
}

func findKthNumber(n int, k int) int {
	curr := 1
	ans := 1
	for curr < k {
		count := getCount(ans, n)
		if curr+count > k {
			ans *= 10
			curr++
		} else {
			ans++
			curr += count
		}
	}
	return ans
}

// @lc code=end

