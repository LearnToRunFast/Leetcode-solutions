/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 *
 * https://leetcode-cn.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (61.04%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    75.4K
 * Total Submissions: 123.3K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
 *
 *
 * 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
 *
 * 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
 * 输出：4
 * 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
 * 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1
 * ，大于 n 的值 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["10", "0", "1"], m = 1, n = 1
 * 输出：2
 * 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 600
 * 1 <= strs[i].length <= 100
 * strs[i] 仅由 '0' 和 '1' 组成
 * 1 <= m, n <= 100
 *
 *
 */

// @lc code=start
// optimized version
func findMaxForm(strs []string, m int, n int) int {
	// define dp[i][j] as the max number of strings that can be formed with i 0s and j 1s
	// dp[i][j]:
	// if i<zeros  ∣  j<ones, then dp[i][j]
	// else, then max(dp[i][j],dp[i-zeros][k−ones]+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros, ones := count(str)
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

// func findMaxForm(strs []string, m int, n int) int {
// 	//define dp[i][j][k] as the max length of the subset of strs[0:i] with j zeros and k ones
// 	// dp[i][j][k]:
// 	// if j<zeros  ∣  k<ones, then dp[i−1][j][k]
// 	// else, then max(dp[i−1][j][k],dp[i−1][j−zeros][k−ones]+1)
// 	dp := make([][][]int, len(strs)+1)
// 	for i := range dp {
// 		dp[i] = make([][]int, m+1)
// 		for j := range dp[i] {
// 			dp[i][j] = make([]int, n+1)
// 		}
// 	}

// 	for i, str := range strs {
// 		zeros, ones := count(str)
// 		for j := 0; j <= m; j++ {
// 			for k := 0; k <= n; k++ {
// 				if j < zeros || k < ones {
// 					dp[i+1][j][k] = dp[i][j][k]
// 				} else {
// 					dp[i+1][j][k] = max(dp[i][j][k], dp[i][j-zeros][k-ones]+1)
// 				}
// 			}
// 		}
// 	}
// 	return dp[len(strs)][m][n]
// }
func count(s string) (int, int) {
	var z, o int
	for _, c := range s {
		if c == '0' {
			z++
		} else {
			o++
		}
	}
	return z, o
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

