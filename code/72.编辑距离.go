/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (61.15%)
 * Likes:    1791
 * Dislikes: 0
 * Total Accepted:    166.9K
 * Total Submissions: 272.3K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * word1 和 word2 由小写英文字母组成
 *
 *
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
		if i == 0 {
			for j := 1; j <= n; j++ {
				dp[0][j] = j
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			insertA := dp[i-1][j] + 1
			insertB := dp[i][j-1] + 1
			changeA := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				changeA++
			}
			dp[i][j] = min(insertA, min(insertB, changeA))
		}
	}
	return dp[m][n]
}

// @lc code=end

