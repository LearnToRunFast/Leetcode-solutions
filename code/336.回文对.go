/*
 * @lc app=leetcode.cn id=336 lang=golang
 *
 * [336] 回文对
 *
 * https://leetcode-cn.com/problems/palindrome-pairs/description/
 *
 * algorithms
 * Hard (40.04%)
 * Likes:    228
 * Dislikes: 0
 * Total Accepted:    20.7K
 * Total Submissions: 51.7K
 * Testcase Example:  '["abcd","dcba","lls","s","sssll"]'
 *
 * 给定一组 互不相同 的单词， 找出所有 不同 的索引对 (i, j)，使得列表中的两个单词， words[i] + words[j]
 * ，可拼接成回文串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["abcd","dcba","lls","s","sssll"]
 * 输出：[[0,1],[1,0],[3,2],[2,4]]
 * 解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：words = ["bat","tab","cat"]
 * 输出：[[0,1],[1,0]]
 * 解释：可拼接成的回文串为 ["battab","tabbat"]
 *
 * 示例 3：
 *
 *
 * 输入：words = ["a",""]
 * 输出：[[0,1],[1,0]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * words[i] 由小写英文字母组成
 *
 *
 */

// @lc code=start
func reverse(s string) string {
	n := len(s)
	var b strings.Builder
	for i := n - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
func isPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
func getPos(dict *map[string]int, s string) int {
	if idx, ok := (*dict)[s]; !ok {
		return -1
	} else {
		return idx
	}
}
func palindromePairs(words []string) [][]int {
	n := len(words)
	dict := map[string]int{}
	rWords := make([]string, n)
	for i, word := range words {
		rWords[i] = reverse(word)
	}
	for i := 0; i < n; i++ {
		dict[rWords[i]] = i
	}
	ans := [][]int{}
	for i, word := range words {
		m := len(word)

		for j := 0; j <= m; j++ {
			if isPalindrome(word[j:m]) {
				pos := getPos(&dict, word[:j])
				if pos != -1 && pos != i {
					ans = append(ans, []int{i, pos})
				}
			}
			if j != 0 && isPalindrome(word[:j]) {
				pos := getPos(&dict, word[j:m])
				if pos != -1 && pos != i {
					ans = append(ans, []int{pos, i})
				}
			}
		}
	}
	return ans

}

// @lc code=end

