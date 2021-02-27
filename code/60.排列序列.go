/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (51.79%)
 * Likes:    475
 * Dislikes: 0
 * Total Accepted:    73.4K
 * Total Submissions: 141.2K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3, k = 3
 * 输出："213"
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 4, k = 9
 * 输出："2314"
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, k = 1
 * 输出："123"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func generateNfactorial(n int) []int {
	table := make([]int, n+1)
	table[0] = 1
	for i := 1; i <= n; i++ {
		table[i] = table[i-1] * i
	}
	return table
}
func getPermutation(n int, k int) string {
	if k < 1 {
		return ""
	}

	lookup := generateNfactorial(n)
	if k > lookup[n] {
		return ""
	}

	sample := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sample[i] = 1
	}
	var ans strings.Builder
	k-- // necesary to make it start from 0

	for i := 1; i <= n; i++ {
		q := k/lookup[n-i] + 1

		for j := 1; j <= n; j++ {
			q -= sample[j]
			if q == 0 {
				sample[j] = 0
				ans.WriteString(strconv.Itoa(j))
				break
			}
		}
		k = k % lookup[n-i]
	}
	return ans.String()
}

// @lc code=end
