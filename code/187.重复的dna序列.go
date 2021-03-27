/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 *
 * https://leetcode-cn.com/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (46.66%)
 * lenikes:    151
 * Dislikes: 0
 * Total Accepted:    30.2K
 * Total Submissions: 64.6K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * 所有 DNA 都由一系列缩写为 'A'，'C'，'G' 和 'T' 的核苷酸组成，例如："ACGAATTCCG"。在研究 DNA 时，识别 DNA
 * 中的重复序列有时会对研究非常有帮助。
 *
 * 编写一个函数来找出所有目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 * 输出：["AAAAACCCCC","CCCCCAAAAA"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "AAAAAAAAAAAAA"
 * 输出：["AAAAAAAAAA"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s[i] 为 'A'、'C'、'G' 或 'T'
 *
 *
 */
package main

// @lc code=start
// func findRepeatedDnaSequences(s string) []string {
// 	n := len(s)
// 	l, r := 0, 10
// 	lookup := map[string]int{}
// 	ans := []string{}
// 	for ; r <= n; r++ {
// 		word := s[l:r]
// 		if v, ok := lookup[word]; ok && v == 1 {
// 			ans = append(ans, word)
// 		}
// 		lookup[word]++
// 		l++
// 	}
// 	return ans
// }
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	len := 10
	ans := []string{}
	if n < 11 {
		return ans
	}

	cToI := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	value := 0
	for i := 0; i < len; i++ {
		// empty lowest 2 bits
		value <<= 2
		// fill up the lowest 2 bits with new value
		value = value | cToI[s[i]]
	}
	lookup := map[int]int{value: 1}
	highestBit := 2 * len
	// keep track of current sub string
	l := 1
	for r := len; r < n; r++ {
		// empty lowest 2 bits
		value <<= 2
		// fill up new value
		value = value | cToI[s[r]]
		// clear old highest 2 bits which at 2 * len
		mask := ^(3 << highestBit)
		value &= mask
		if v := lookup[value]; v == 1 {
			ans = append(ans, s[l:r+1])
		}
		l++
		lookup[value]++
	}
	return ans
}

// @lc code=end
