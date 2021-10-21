/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找共用字符
 *
 * https://leetcode-cn.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (73.58%)
 * Likes:    242
 * Dislikes: 0
 * Total Accepted:    57.5K
 * Total Submissions: 78.2K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * 给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序
 * 返回答案。
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["bella","label","roller"]
 * 输出：["e","l","l"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：words = ["cool","lock","cook"]
 * 输出：["c","o"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 100
 * 1 <= words[i].length <= 100
 * words[i] 由小写英文字母组成
 *
 *
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func commonChars(words []string) []string {
	n := len(words)
	lookup := make([][]int, n)
	for i, word := range words {
		lookup[i] = make([]int, 26)
		for _, c := range word {
			lookup[i][c-'a']++
		}
	}
	ans := make([]string, 0)
	for i := 0; i < 26; i++ {
		val := lookup[0][i]
		for _, arr := range lookup {
			val = min(val, arr[i])
		}
		str := string(i + 'a')
		for j := 0; j < val; j++ {
			ans = append(ans, str)
		}
	}

	return ans
}

// @lc code=end

