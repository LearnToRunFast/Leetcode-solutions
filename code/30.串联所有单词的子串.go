/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (33.38%)
 * Likes:    417
 * Dislikes: 0
 * Total Accepted:    56.6K
 * Total Submissions: 167.1K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 *
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 *
 *
 * 示例 2：
 *
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 *
 *
 */
package main

// @lc code=start
// func findSubstring(s string, words []string) []int {
// 	lookup := map[string]int{}

// 	for _, v := range words {
// 		lookup[v]++
// 	}

// 	wordLen := len(words[0])
// 	stringLen := wordLen * len(words)
// 	ans := []int{}

// 	for i := 0; i <= len(s)-stringLen; i++ {
// 		window := map[string]int{}
// 		count := 0
// 		end := i + stringLen

// 		for j := i; j < end; j += wordLen {
// 			wordEnd := j + wordLen
// 			window[s[j:wordEnd]]++
// 			if _, ok := lookup[s[j:wordEnd]]; !ok || lookup[s[j:wordEnd]] < window[s[j:wordEnd]] {
// 				break
// 			}
// 			count++
// 		}
// 		if count == len(words) {
// 			ans = append(ans, i)
// 		}

// 	}
// 	return ans
// }
func findSubstring(s string, words []string) []int {
	lookup := map[string]int{}

	for _, v := range words {
		lookup[v]++
	}

	wordLen := len(words[0])
	wordSize := len(words)
	stringLen := len(s)
	ans := []int{}

	for i := 0; i < wordLen; i++ {
		window := map[string]int{}
		left, right, count := i, i, 0
		for right+wordLen <= stringLen {
			word := s[right : right+wordLen]
			right += wordLen
			if _, ok := lookup[word]; !ok {
				count = 0
				left = right
				window = map[string]int{}
				continue
			}

			window[word]++
			count++
			for window[word] > lookup[word] {
				count--
				window[s[left:left+wordLen]]--
				left += wordLen
			}
			if count == wordSize {
				ans = append(ans, left)
			}
		}

	}
	return ans
}

// @lc code=end
