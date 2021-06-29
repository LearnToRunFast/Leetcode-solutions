/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 *
 * https://leetcode-cn.com/problems/longest-repeating-character-replacement/description/
 *
 * algorithms
 * Medium (52.85%)
 * Likes:    457
 * Dislikes: 0
 * Total Accepted:    46.7K
 * Total Submissions: 88.4K
 * Testcase Example:  '"ABAB"\n2'
 *
 * 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k
 * 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
 *
 * 注意：字符串长度 和 k 不会超过 10^4。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ABAB", k = 2
 * 输出：4
 * 解释：用两个'A'替换为两个'B',反之亦然。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "AABABBA", k = 1
 * 输出：4
 * 解释：
 * 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
 * 子串 "BBBB" 有最长重复字母, 答案为 4。
 *
 *
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	dict := [26]int{}
	l, r := 0, 0
	max := 0
	n := len(s)
	for ; r < n; r++ {
		dict[s[r]-'A']++
		if max < dict[s[r]-'A'] {
			max = dict[s[r]-'A']
		}
		if r-l+1-max > k {
			dict[s[l]-'A']--
			l++
		}
	}
	return n - l
}

// @lc code=end

