/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (66.78%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    15.1K
 * Total Submissions: 22.6K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * 给定一个字符串数组 words，找到 length(word[i]) * length(word[j])
 * 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
 *
 * 示例 1:
 *
 * 输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * 输出: 16
 * 解释: 这两个单词为 "abcw", "xtfn"。
 *
 * 示例 2:
 *
 * 输入: ["a","ab","abc","d","cd","bcd","abcd"]
 * 输出: 4
 * 解释: 这两个单词为 "ab", "cd"。
 *
 * 示例 3:
 *
 * 输入: ["a","aa","aaa","aaaa"]
 * 输出: 0
 * 解释: 不存在这样的两个单词。
 *
 */

// @lc code=start
func maxProduct(words []string) int {
	n := len(words)
	dict := make([]int, n)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) <= len(words[j])
	})
	for i, word := range words {
		for _, c := range word {
			dict[i] |= 1 << (c - 'a')
		}
	}
	ans := 0
	end := 0
	for i := n - 1; i >= 0; i-- {
		m := len(words[i])
		start := i - 1
		if m*m <= ans || start < end {
			break
		}
		for j := start; j >= end; j-- {
			if dict[i]&dict[j] == 0 {
				newVal := m * len(words[j])
				if ans < newVal {
					ans = newVal
					end = j
				}
			}
		}
	}

	return ans

}

// @lc code=end

